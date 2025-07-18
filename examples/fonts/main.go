// This works in the web too!
package main

import (
	"embed"
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

//go:embed font.ttf
var embeddedFiles embed.FS

func main() {
	w, err := vuelto.NewWindow("hi", 800, 600, false, false)
	if err != nil {
		log.Fatalln(err)
	}

	ren := w.NewRenderer2D()
	uir := w.NewUIRenderer()

	font, err := uir.LoadFont(vuelto.FontEmbed{Filesystem: embeddedFiles, Font: "font.ttf"}, "Hello, world!", -0.1, -0.1, 30)
	if err != nil {
		log.Fatalln(err)
	}

	for !w.Close() {
		ren.ClearColor([4]int{100, 100, 100, 255})

		font.Draw()

		w.Refresh()
	}
}
