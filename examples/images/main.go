// This works in the web too! This is because of the images being embedded!
package main

import (
	"embed"
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

//go:embed tree.png galaxy.png
var embeddedFiles embed.FS

func main() {
	w, err := vuelto.NewWindow("Image Example - Vuelto", 800, 600, true, false)
	if err != nil {
		log.Fatalln(err)
	}
	ren := w.NewRenderer2D()

	imageEmbedOne := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "tree.png",
	}

	imageEmbedTwo := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "galaxy.png",
	}

	imageOne, err := ren.LoadImage(imageEmbedOne, 0.7, 0.3, -0.5, 0.5, nil)
	if err != nil {
		log.Fatalln(err)
	}

	imageTwo, err := ren.LoadImage(imageEmbedTwo, 0, 0, 1, 1, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for !w.Close() {
		imageTwo.Draw()
		imageOne.Draw()
		w.Refresh()
	}
}
