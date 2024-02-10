package main

import (
	"github.com/vuelto-org/vuelto"
)

func main() {
	a := vuelto.NewApp()

	w1 := a.NewWindow("hi", 800, 600, false)
	w2 := a.NewWindow("hi2", 800, 600, false)

	ren := a.NewRenderer2D()

	image := ren.LoadImage("test/image.png", 300, 300, 250, 250)
	image1 := ren.LoadImage("test/image.png", 100, 100, 150, 150)

	for !w1.Close() && !w2.Close() {
		w1.SetContextCurrent()
		ren.ClearColor([3]float32{0.3, 0.4, 0.3})

		ren.DrawRect(0, 0, 500, 500, [3]float32{0.1, 0.5, 0.7})

		w1.Refresh()
		w2.SetContextCurrent()

		image.Draw()
		image1.Draw()

		w2.Refresh()

	}
}
