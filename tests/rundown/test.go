package main

import (
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w1, err := vuelto.NewWindow("hi", 800, 600, false, false)
	if err != nil {
		log.Fatalln(err)
	}

	w2, err := vuelto.NewWindow("hi2", 800, 600, false, false)
	if err != nil {
		log.Fatalln(err)
	}

	ren1 := w1.NewRenderer2D()
	ren2 := w2.NewRenderer2D()

	rect, err := ren1.NewRect(0, 0, 0.5, 0.5, [4]int{10, 145, 245, 255})
	if err != nil {
		log.Fatalln(err)
	}

	line, err := ren1.NewLine(0.1, 0.1, 0.4, 0.4, [4]int{10, 145, 245, 255})
	if err != nil {
		log.Fatalln(err)
	}

	pixelmap, err := ren1.NewPixelmap()
	if err != nil {
		log.Fatalln(err)
	}

	image, err := ren2.LoadImage("test/test/tree.png", 0.1, -0.1, 0.4, -0.4, nil)
	if err != nil {
		log.Fatalln(err)
	}

	image1, err := ren2.LoadImage("test/test/galaxy.png", -0.1, -0.1, 0.4, 0.4, nil)
	if err != nil {
		log.Fatalln(err)
	}

	image2, err := ren2.LoadImage(vuelto.ImageHTTP{
		Url: "https://dev-tester.com/content/images/2021/12/blog_cover_further_api_testing_with_http_toolkit.png",
	}, -0.1, 0.1, 0.4, 0.4, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for !w1.Close() && !w2.Close() {

		ren1.ClearColor([4]int{100, 100, 100, 255})

		if w1.KeyPressed(vuelto.Keys["Left"]) {
			rect.Pos.X = rect.Pos.X - 0.5*w1.GetDeltaTime()
		} else if w1.KeyPressed(vuelto.Keys["Right"]) {
			rect.Pos.X = rect.Pos.X + 0.5*w1.GetDeltaTime()
		}

		pixelmap.SetPixel(200, 500, [4]int{34, 28, 128, 255})

		log.Println(w1.MousePos())

		rect.Draw()

		line.Draw()
		image2.Draw()

		image.Draw()
		image1.Draw()

		pixelmap.Draw()

		w1.Refresh()
		w2.Refresh()
	}
}
