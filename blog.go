package main

import (
	"html/template"
	"time"
	"regexp"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"sync"
	"bufio"
)

type BlogConfig struct {
	SrcDir  string
	SiteDir string
}

var config = BlogConfig{
	SrcDir:  "src",
	SiteDir: "_site",
}

// -----------------------------------------------------------------------------
type Blog struct {
	Config      BlogConfig
	Posts       []Post
	BuildErrors chan error
	BuildLocks  map[string]sync.Mutex
}

func NewBlog(config BlogConfig) Blog {
	return Blog{
		Config:      config,
		BuildErrors: make(chan error),
	}
}

func (blog Blog) PostSrcDir() string {
	return filepath.Join(blog.Config.SrcDir, "_posts")
}

func (blog *Blog) LoadPosts() error {
	fileInfos, err := ioutil.ReadDir(filepath.Join(blog.Config.SrcDir, "_posts"))

	if err != nil {
		return nil
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		if err := blog.LoadPost(fileInfo.Name()); err != nil {
			return err
		}
	}

	return nil
}

// -----------------------------------------------------------------------------
func (blog Blog) BuildPosts() error {
	if len(blog.Posts) == 0 {
		return nil
	}

	if err := os.RemoveAll(filepath.Join(blog.Config.SiteDir, "posts")); err != nil {
		return err
	}

	for _, post := range blog.Posts {
		layout := fmt.Sprintf("%s.gohtml", post.Layout())
		t := template.New(layout)
		layoutPath := filepath.Join(blog.Config.SrcDir, "_layouts", layout)
		t, err := t.ParseFiles(layoutPath)
		if err != nil {
			return err
		}

		go func() {
			destPath := filepath.Join(blog.Config.SiteDir, "posts", post.DestPath)

			dir, _ := filepath.Split(destPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				blog.BuildErrors <- err
				return
			}

			dest, err := os.Create(destPath)
			if err != nil {
				blog.BuildErrors <- err
				return
			}
			defer dest.Close()

			if err = t.Execute(dest, post); err != nil {
				blog.BuildErrors <- err
				return
			}

			blog.BuildErrors <- nil
		}()
	}

	for i := 0; i < len(blog.Posts); i++ {
		if err := <-blog.BuildErrors; err != nil {
			return err
		}
	}

	return nil
}

// -----------------------------------------------------------------------------
type Post struct {
	FrontMatter map[string]string
	Filename    string
	DestPath    string
	Date        time.Time
	Content     template.HTML
}

func (post Post) Layout() string {
	if l, exists := post.FrontMatter["layout"]; exists {
		return l
	}

	return "default"
}

var postRegex = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})-(.+).md`)
var frontMatterRegex = regexp.MustCompile(`^([a-zA-Z0-9]+): (.+[^\s])`)

func (blog *Blog) LoadPost(filename string) error {

	if ! postRegex.MatchString(filename) {
		return fmt.Errorf("invalid post filename '%s'", filename)
	}

	matches := postRegex.FindStringSubmatch(filename)[1:]

	date, err := time.Parse("2006-01-02", matches[0])
	if err != nil {
		return fmt.Errorf("invalid post date '%s'", matches[0])
	}

	destPath := fmt.Sprintf("%s/%s.html", date.Format("2006/01/02"), matches[1])

	f, err := os.Open(filepath.Join(blog.Config.SrcDir, "_posts", filename))
	if err != nil {
		return err
	}
	defer f.Close()

	var md []byte
	fm := make(map[string]string)
	reader := bufio.NewReader(f)
	state := "0"

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		switch state {
		case "0":
			if string(line) == "---\n" {
				state = "f"
				continue
			} else {
				state = "p"
				md = append(md, line...)
			}
		case "f":
			if string(line) == "---\n" {
				state = "p"
				continue
			} else {
				if ! frontMatterRegex.Match(line) {
					return fmt.Errorf("invalid front matter '%s'", string(line))
				}

				matches := frontMatterRegex.FindSubmatch(line)[1:]
				fm[string(matches[0])] = string(matches[1])
			}
		case "p":
			md = append(md, line...)
		}
	}

	blog.Posts = append(blog.Posts, Post{
		FrontMatter: fm,
		Filename: filename,
		Date:     date,
		DestPath: destPath,
		Content:  Render(md),
	})

	return nil
}
