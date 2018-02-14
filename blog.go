package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// -----------------------------------------------------------------------------
type Blog struct {
	Config BlogConfig
	Posts  []Document
}

func NewBlog(config BlogConfig) Blog {
	if err := os.RemoveAll(config.SiteDir); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(config.SiteDir, 0755); err != nil {
		panic(err)
	}

	return Blog{
		Config: config,
		Posts:  make([]Document, 0),
	}
}

func (blog *Blog) LoadPosts(srcDir string) error {
	srcDir = filepath.Join(blog.Config.SrcDir, srcDir)
	fileInfos, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			continue
		}

		doc, err := blog.LoadDocument(filepath.Join(srcDir, info.Name()))
		if err != nil {
			return err
		}

		if doc.FrontMatter.Layout == "" {
			doc.FrontMatter.Layout = "default"
		}

		blog.Posts = append(blog.Posts, doc)
	}

	sort.Slice(blog.Posts, func(i, j int) bool {
		return blog.Posts[i].FrontMatter.Date.After(blog.Posts[j].FrontMatter.Date)
	})

	return nil
}

type TemplateData struct {
	Blog Blog
	Doc  Document
}

func (blog Blog) Build() error {
	err := filepath.Walk(blog.Config.SrcDir, func(path string, info os.FileInfo, err error) error {
		if info.Name()[0] == '_' {
			fmt.Printf("Skipping private directory %s\n", info.Name())
			return filepath.SkipDir
		}

		if info.IsDir() {
			if info.Name() == "posts" {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.Mode().IsRegular() {
			fmt.Printf("Seeing irregular %s\n", info.Name())
			return nil
		}

		fmt.Printf("Seeing file %s\n", info.Name())
		switch filepath.Ext(path) {
		case ".gohtml", ".md", ".html":
			doc, err := blog.LoadDocument(path)
			if err != nil {
				return err
			}

			if err := blog.BuildDocument(doc); err != nil {
				return err
			}
		default:
			if err := blog.Copy(path); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (blog Blog) Copy(path string) error {
	src, err := os.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	destPath := filepath.Join(
		config.SiteDir, path[len(config.SrcDir):],
	)

	// make the destination directory in case it does not exist
	destDir, _ := filepath.Split(destPath)
	if _, err := os.Stat(destDir); err != nil {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return err
		}
	}

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	return dest.Sync()
}

var templateFuncMap = template.FuncMap{
	"absurl": func(path string) string {
		u, err := url.Parse(config.BaseURL)
		if err != nil {
			panic(err)
		}
		u.Path = filepath.Join(u.Path, path)
		return u.String()
	},
	"fdate": func(date time.Time, layout string) string {
		return date.Format(layout)
	},
}

func (blog Blog) BuildDocument(doc Document) error {
	destPath := filepath.Join(
		config.SiteDir,
		doc.Path[:len(doc.Path)-len(doc.Ext)]+".html",
	)

	destDir, _ := filepath.Split(destPath)

	// make the destination directory in case it does not exist
	if _, err := os.Stat(destDir); err != nil {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return err
		}
	}

	// create the dest file
	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	if doc.FrontMatter.Layout != "" {
		layoutName := doc.FrontMatter.Layout + ".gohtml"
		layoutPath := filepath.Join(blog.Config.SrcDir, "_layouts", layoutName)

		t, err := template.New(layoutName).Funcs(templateFuncMap).ParseFiles(layoutPath)
		if err != nil {
			return err
		}

		if err := t.Execute(dest, TemplateData{blog, doc}); err != nil {
			return err
		}
	} else {

		_, name := filepath.Split(doc.Path)
		t, err := template.New(name).Funcs(templateFuncMap).ParseFiles(filepath.Join(blog.Config.SrcDir, doc.Path))
		if err != nil {
			return err
		}

		if err := t.Execute(dest, TemplateData{Blog: blog}); err != nil {
			return err
		}
	}

	return dest.Sync()
}
