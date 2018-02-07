package main

import (
	"testing"
	"html/template"
	"os"
	"io/ioutil"
	"path/filepath"
)

var testConfig = BlogConfig{
	SrcDir:  "test_src",
	SiteDir: "test_site",
}

// -----------------------------------------------------------------------------
func TestBlog_LoadPosts(t *testing.T) {
	blog := NewBlog(testConfig)

	err := blog.LoadPosts()
	if err != nil {
		t.Fatal(err)
	}

	if len(blog.Posts) != 1 {
		t.Fatalf("expected len posts to be 1, saw %d", len(blog.Posts))
	}

	post := blog.Posts[0]

	filename := "2018-02-01-example-post.md"
	if post.Filename != filename {
		t.Fatalf("%s != %s", filename, post.Filename)
	}

	if result := post.Date.Format("02-Jan-2006"); result != "01-Feb-2018" {
		t.Fatalf("%s != %s", "01-Feb-2018", result)
	}

	if result := post.DestPath; result != "2018/02/01/example-post.html" {
		t.Fatalf("%s != %s", "2018/02/01/example-post.html", result)
	}

	if result := post.Layout(); result != "posts" {
		t.Fatalf("%s != %s", "posts", result)
	}

	if result := post.Title(); result != "Example Post" {
		t.Fatalf("%s != %s", "Example Post", result)
	}

	expectedContent := "<h1>This is a post</h1>\n\n"
	expectedContent += "<p>Hello there</p>\n"

	if result := post.Content; result != template.HTML(expectedContent) {
		t.Fatalf("%q != %q", expectedContent, result)
	}
}

// -----------------------------------------------------------------------------
func TestBlog_BuildPosts(t *testing.T) {
	os.RemoveAll(testConfig.SiteDir)
	os.MkdirAll(testConfig.SiteDir, 0755)

	blog := NewBlog(testConfig)

	blog.LoadPosts()
	if err := blog.BuildPosts(); err != nil {
		t.Fatal(err)
	}


	output, err := ioutil.ReadFile(
		filepath.Join(blog.Config.SiteDir, "posts/2018/02/01/example-post.html"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "<!doctype html>" + "\n"
	expected += "<html lang=\"en\">" + "\n"
	expected += "<head>" + "\n"
	expected += "  <meta charset=\"UTF-8\">" + "\n"
	expected += "  <title>Posts</title>" + "\n"
	expected += "</head>" + "\n"
	expected += "<body>" + "\n"
	expected += "<h1>This is a post</h1>" + "\n\n"
	expected += "<p>Hello there</p>" + "\n\n"
	expected += "</body>" + "\n"
	expected += "</html>" + "\n"

	if result := output; string(result) != expected {
		t.Fatalf("\n%q\n!=\n%q\n", expected, string(result))
	}
}
