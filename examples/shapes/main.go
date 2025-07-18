// This works in the web too!
package main

import (
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w, err := vuelto.NewWindow("hi", 800, 600, true, false)
	if err != nil {
		log.Fatalln(err)
	}

	ren := w.NewRenderer2D()

	rect, err := ren.NewRect(0, 0, -1, -1, [4]int{10, 145, 245, 255})
	if err != nil {
		log.Fatalln(err)
	}

	rect2, err := ren.NewRect(0, 0, 1, 1, [4]int{245, 145, 10, 255})
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

		w.Refresh()
	}
}
