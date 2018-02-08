package main

import (
	"io"
	"regexp"
	"fmt"
	"bufio"
	"errors"
	"strings"
)

type Document struct {
	FrontMatter map[string]string
	Content     []byte
}

var frontMatterDelimRegex = regexp.MustCompile(`^---\s*$`)
var frontMatterRegex = regexp.MustCompile(`^([a-zA-Z0-9]+): (.+[^\s])`)

const (
	Parse0 = iota
	Parse1
	Parse2
	Parse3
	Parse4
	Parse5
)

func ReadDocument(r io.Reader) (Document, error) {
	reader := bufio.NewReader(r)

	doc := Document{
		FrontMatter: make(map[string]string),
		Content:     make([]byte, 0),
	}

	state := Parse0
	for state != Parse5 {
		input, err := reader.ReadBytes('\n')
		if len(input) == 0 && err != nil {
			input = nil
		}

		switch state {
		case Parse0:
			if frontMatterDelimRegex.Match(input) {
				state = Parse1
			} else if input != nil {
				state = Parse4
			} else {
				return Document{}, errors.New("empty document")
			}
		case Parse1, Parse2:
			if frontMatterRegex.Match(input) {
				state = Parse2
			} else if frontMatterDelimRegex.Match(input) {
				state = Parse3
			} else {
				return Document{}, fmt.Errorf("invalid front matter '%s'", input)
			}
		case Parse3:
			if input != nil {
				state = Parse4
			} else {
				return Document{}, errors.New("empty document")
			}
		case Parse4:
			if input == nil {
				state = Parse5
			}
		default:
			panic(fmt.Sprintf("invalid state %d", state))
		}

		if state == Parse2 {
			matches := frontMatterRegex.FindSubmatch(input)[1:]
			doc.FrontMatter[strings.ToLower(string(matches[0]))] = string(matches[1])
		} else if state == Parse4 {
			doc.Content = append(doc.Content, input...)
		}
	}

	return doc, nil
}
