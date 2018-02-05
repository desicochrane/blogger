package main

import (
	"github.com/Depado/bfchroma"
	"html/template"
	"github.com/alecthomas/chroma/formatters/html"
	"gopkg.in/russross/blackfriday.v2"
)

func Render(md []byte) template.HTML {
	return template.HTML(blackfriday.Run(
		md,
		blackfriday.WithRenderer(
			bfchroma.NewRenderer(
				bfchroma.WithoutAutodetect(),
				bfchroma.ChromaOptions(
					html.WithLineNumbers(),
					html.WithClasses(),
				),
			),
		),
	))
}
