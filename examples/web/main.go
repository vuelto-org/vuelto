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

	imageOne, err := ren.LoadImage(imageEmbedOne, 0.5, 0.5, -0.5, 0.5, nil)
	if err != nil {
		log.Fatalln(err)
	}

	imageTwo, err := ren.LoadImage(imageEmbedTwo, 0, 0, 1, 1, nil)
	if err != nil {
		log.Fatalln(err)
	}

	imageThree, err := ren.LoadImage(vuelto.ImageHTTP{
		Url: "https://dev-tester.com/content/images/2021/12/blog_cover_further_api_testing_with_http_toolkit.png",
	}, -0.1, 0.1, 0.4, 0.4, nil)
	if err != nil {
		log.Fatalln(err)
	}

	rect, err := ren.NewRect(0, 0, -1, -1, [4]int{10, 145, 245, 255})
	if err != nil {
		log.Fatalln(err)
	}

	rect2, err := ren.NewRect(0, 1, 1, 1, [4]int{245, 145, 10, 255})
	if err != nil {
		log.Fatalln(err)
	}

	line, err := ren.NewLine(0.5, 0.5, -0.5, -0.5, [4]int{10, 145, 245, 255})
	if err != nil {
		log.Fatalln(err)
	}

	for !w.Close() {
		ren.ClearColor([4]int{100, 100, 100, 255})
		ren.DrawLine(-0.9, -0.9, 0.9, -0.9, [4]int{10, 145, 245, 255})

		rect.Draw()
		rect2.Draw()
		line.Draw()

		imageOne.Draw()
		imageTwo.Draw()
		imageThree.Draw()
		w.Refresh()
	}
}
