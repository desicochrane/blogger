package main

import (
	"testing"
	"strings"
)

var eol = "\n"

func TestReadDocument(t *testing.T) {
	frontMatter := `---
Title: This is my lovely post 
date: 2018-02-02
layout: post
---
`

	content := `# Lovely post

Let \(b_{baby}(n)\) and \(b_{adult}(n)\) denote the number of 
baby and adult bunnies:

$$ F(n) = b_{adult}(n) + b_{baby}(n) $$
`

	doc, err := ReadDocument(strings.NewReader(frontMatter + content))
	if err != nil {
		t.Fatal(err)
	}

	samples := [][]string{
		{"title", "This is my lovely post"},
		{"date", "2018-02-02"},
		{"layout", "post"},
	}

	for _, e := range samples {
		if doc.FrontMatter[e[0]] != e[1] {
			t.Fatalf("%s != %s", doc.FrontMatter[e[0]], e[1])
		}
	}

	if string(doc.Content) != content {
		t.Fatalf("\n%q\n!=\n%q", content, string(doc.Content))
	}
}
