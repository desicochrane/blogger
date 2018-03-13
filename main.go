package main

import (
	"flag"
	"net/http"

	"fmt"
	"log"

	"os"

	"strings"

	"github.com/joho/godotenv"
)

type BlogConfig struct {
	BaseURL string
	SrcDir  string
	SiteDir string
}

var config = BlogConfig{
	BaseURL: "http://localhost",
	SrcDir:  "src",
	SiteDir: "docs",
}

func main() {
	serve := flag.Bool("serve", false, "serve blog")
	port := flag.String("port", "1234", "http port")
	env := flag.String("env", "dev", "build mode")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if baseURL := os.Getenv(strings.ToUpper(*env) + "_URL"); baseURL != "" {
		config.BaseURL = baseURL
	}

	build(*env)

	if *serve {
		http.Handle("/", http.FileServer(http.Dir(config.SiteDir)))
		fmt.Printf("\nServing on port :%s\n", *port)
		http.ListenAndServe(":" + *port, nil)
	}
}

func build(env string) {
	blog := NewBlog(config)

	if err := blog.LoadPosts("posts"); err != nil {
		panic(err)
	}

	if strings.ToUpper(env) == "DEV" {
		if err := blog.LoadPosts("_drafts"); err != nil {
			panic(err)
		}
	}

	for _, doc := range blog.Posts {
		if err := blog.BuildDocument(doc); err != nil {
			panic(err)
		}
	}

	if err := blog.Build(); err != nil {
		panic(err)
	}
}
