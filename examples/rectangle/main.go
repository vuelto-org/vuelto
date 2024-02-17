package main

import (
	vuelto "github.com/vuelto-org/vuelto/pkg"
)

func main() {
	w := vuelto.NewWindow("hi", 800, 600, false)

	ren := w.NewRenderer2D()

	for !w.Close() {
		ren.ClearColor([3]float32{0.3, 0.4, 0.3})

		ren.DrawRect(0, 0, 500, 500, [3]float32{0.1, 0.5, 0.7})

		w.Refresh()

	}
}
