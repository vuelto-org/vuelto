// Warning: This example is not working in the web! This is due to the web not supporting multiple windows and due to the image loading.
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

	image, err := ren1.LoadImage("examples/two-windows/galaxy.png", 0, 0, 0.5, 0.5, nil)
	if err != nil {
		log.Fatalln(err)
	}

	image1, err := ren2.LoadImage("examples/two-windows/tree.png", -0.5, -0.5, 0.5, 0.5, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for !w1.Close() && !w2.Close() {
		ren1.ClearColor([4]int{100, 100, 100, 255})
		ren2.ClearColor([4]int{100, 100, 100, 255})

		ren1.DrawRect(-0.7, 0.7, 0.7, 0.7, [4]int{10, 145, 245, 255})

		image.Draw()
		image1.Draw()

		w1.Refresh()
		w2.Refresh()
	}
}
