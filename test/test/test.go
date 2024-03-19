package main

import (
	"log"

	"vuelto.me/pkg"
)

func main() {
	w1, err1 := vuelto.NewWindow("hi", 800, 600, false)
  if err1 != nil {
    log.Fatalln("Could not create a new window1: ", err1)
  }
	w2, err2 := vuelto.NewWindow("hi2", 800, 600, false)
  if err2 != nil {
    log.Fatalln("Could not create a new window1: ", err1)
  }

	ren1 := w1.NewRenderer2D()
	ren2 := w2.NewRenderer2D()

	image := ren2.LoadImage("test/test/tree.png", 300, 300, 250, 250)
	image1 := ren2.LoadImage("test/test/image.png", 100, 100, 150, 150)

	for !w1.Close() && !w2.Close() {
		w1.SetContextCurrent()
		ren1.ClearColor([4]int{100, 100, 100, 255})

		ren1.DrawRect(0, 0, 500, 500, [4]int{10, 145, 245, 255})

		w1.Refresh()
		w2.SetContextCurrent()

		image.Draw()
		image1.Draw()

		w2.Refresh()

	}
}
