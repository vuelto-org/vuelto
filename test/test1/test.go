package main

import (
  "github.com/dimkauzh/vuelto"
)

func main() {
  a := vuelto.NewApp()
  w := a.NewWindow("hi", 800, 600, false)

  ren := a.NewRenderer2D()

  image := ren.LoadImage("test/image.png", 300, 300, 250, 250)
  image1 := ren.LoadImage("test/image.png", 100, 100, 150, 150)

  for !w.Close() {
    ren.ClearColor([3]float32{0.3, 0.4, 0.3})

    ren.DrawRect(0, 0, 500, 500, [3]float32{0.1, 0.5, 0.7})

    image.Draw()
    image1.Draw()

    w.Refresh()
  }
}

