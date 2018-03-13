package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	frontMatterDelimRegex = regexp.MustCompile(`^---\s*$`)
	frontMatterRegex      = regexp.MustCompile(`^([a-zA-Z0-9]+): (.+[^\s])`)
)

const (
	DocStart    = iota
	DocFMStart
	DocFMKeyVal
	DocFMEnd
	DocContent
	DocEnd
)

type FrontMatter struct {
	Title   string
	Date    time.Time
	Layout  string
	Summary string
	Vars    map[string]string
}

func (fm *FrontMatter) set(key string, value string) error {
	switch strings.ToLower(key) {
	case "title":
		fm.Title = value
	case "layout":
		fm.Layout = value
	case "date":
		date, err := time.Parse("02-Jan-2006", value)
		if err != nil {
			return err
		}
		fm.Date = date
	default:
		fm.Vars[key] = value
	}

	return nil
}

type Document struct {
	URL         string
	Path        string
	Ext         string
	FrontMatter FrontMatter
	Content     template.HTML
}

func (blog Blog) LoadDocument(path string) (Document, error) {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return Document{}, err
	}
	defer f.Close()

	ext := filepath.Ext(path)
	fm, content, err := parseDocContent(f, ext)
	if err != nil {
		return Document{}, err
	}

	newExtension := ".html"
	if ext == ".goxml" {
		newExtension = ".xml"
	}
	return Document{
		URL:         path[len(blog.Config.SrcDir):len(path)-len(ext)] + newExtension,
		Path:        path[len(blog.Config.SrcDir):],
		Ext:         ext,
		FrontMatter: fm,
		Content:     content,
	}, nil
}

func parseDocContent(r io.Reader, ext string) (fm FrontMatter, content template.HTML, err error) {
	reader := bufio.NewReader(r)

	rawContent := make([]byte, 0)

	state := DocStart
	for state != DocEnd {
		input, err := reader.ReadBytes('\n')
		if len(input) == 0 && err != nil {
			input = nil
		}

		switch state {
		case DocStart:
			if frontMatterDelimRegex.Match(input) {
				state = DocFMStart
			} else if input != nil {
				state = DocContent
			} else {
				return FrontMatter{}, "", errors.New("empty document")
			}
		case DocFMStart, DocFMKeyVal:
			if frontMatterRegex.Match(input) {
				state = DocFMKeyVal
			} else if frontMatterDelimRegex.Match(input) {
				state = DocFMEnd
			} else {
				return FrontMatter{}, "", fmt.Errorf("invalid front matter '%s'", input)
			}
		case DocFMEnd:
			if input != nil {
				state = DocContent
			} else {
				return FrontMatter{}, "", errors.New("empty document")
			}
		case DocContent:
			if input == nil {
				state = DocEnd
			}
		default:
			panic(fmt.Sprintf("invalid state %d", state))
		}

		if state == DocFMKeyVal {
			matches := frontMatterRegex.FindSubmatch(input)[1:]
			if err := fm.set(string(matches[0]), string(matches[1])); err != nil {
				return FrontMatter{}, "", err
			}
		} else if state == DocContent {
			rawContent = append(rawContent, input...)
		}
	}

	if ext == ".md" {
		content = Render(rawContent)
	} else {
		content = template.HTML(rawContent)
	}

	return fm, content, nil
}
