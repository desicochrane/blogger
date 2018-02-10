package main

import (
	"testing"
)

func TestBlog_LoadPosts(t *testing.T) {
	blog := NewBlog(BlogConfig{
		BaseURL: "https://des.io/blog",
		SrcDir:  "test_src",
		SiteDir: "test_docs",
	})

	blog.LoadPosts("posts")

	expected := map[string][]string{
		"/posts/2018-02-01-example-post.md": {
			"This is my lovely post",
			"2018-02-01",
			"post",
			"<h1>Lovely post</h1>\n",
			"https://des.io/blog/posts/2018-02-01-example-post.html",
		},
	}

	for i := 0; i < len(expected); i++ {
		doc := blog.Posts[i]

		e, found := expected[doc.Path]
		if !found {
			t.Fatalf("slug not found %s", doc.Path)
		}

		if e[0] != doc.FrontMatter.Title {
			t.Fatalf("%s != %s", e[0], doc.FrontMatter.Title)
		}

		if r := doc.FrontMatter.Date.Format("2006-01-02"); e[1] != r {
			t.Fatalf("%s != %s", e[1], r)
		}

		if r := doc.FrontMatter.Layout; e[2] != r {
			t.Fatalf("%s != %s", e[2], r)
		}

		if r := string(doc.Content); e[3] != r {
			t.Fatalf("%s != %s", e[3], r)
		}

		if r := doc.URL; e[4] != r {
			t.Fatalf("%s != %s", e[4], r)
		}
	}
}
