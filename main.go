package main

func main() {
	if err := Build(); err != nil {
		panic(err)
	}

	blog := NewBlog(config)

	if err := blog.LoadPosts(); err != nil {
		panic(err)
	}

	if err := blog.BuildPosts(); err != nil {
		panic(err)
	}

	//http.Handle("/", http.FileServer(http.Dir("_site")))
	//
	//fmt.Printf("Serving on port %s\n", ":1212")
	//http.ListenAndServe(":1212", nil)
}
