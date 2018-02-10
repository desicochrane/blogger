package main

import (
	"html/template"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
)

func Render(md []byte) template.HTML {
	return template.HTML(blackfriday.Run(
		md,
		blackfriday.WithRenderer(
			bfchroma.NewRenderer(
				bfchroma.WithoutAutodetect(),
				bfchroma.ChromaOptions(
					html.WithClasses(),
				),
			),
		),
	))
}
