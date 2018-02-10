package main

import (
	"testing"
)

func TestLoadDocument(t *testing.T) {
	blog := NewBlog(BlogConfig{
		SrcDir:  "test_src",
		SiteDir: "test_docs",
	})
	doc, err := blog.LoadDocument("test_src/posts/2018-02-01-example-post.md")
	if err != nil {
		t.Fatal(err)
	}

	if doc.Path != "/posts/2018-02-01-example-post.md" {
		t.Fatalf("%s != %s", "/posts/2018-02-01-example-post.md", doc.Path)
	}

	if doc.FrontMatter.Title != "This is my lovely post" {
		t.Fatalf("%s != %s", "This is my lovely post", doc.FrontMatter.Title)
	}

	if doc.FrontMatter.Date.Format("02-Jan-2006") != "01-Feb-2018" {
		t.Fatalf("%s != %s", "01-Feb-2018", doc.FrontMatter.Date.Format("02-Jan-2006"))
	}

	if doc.FrontMatter.Layout != "post" {
		t.Fatalf("%s != %s", "post", doc.FrontMatter.Layout)
	}

	if string(doc.Content) != "<h1>Lovely post</h1>\n" {
		t.Fatalf("%q != %q", "<h1>Lovely post</h1>\n", string(doc.Content))
	}
}
