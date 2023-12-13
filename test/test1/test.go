package main

import "github.com/dimkauzh/vuelto"

func main() {
  a := vuelto.NewApp()
  w := a.NewWindow("hi", 800, 600, false)

  ren := w.NewRenderer2D()

  for w.ShouldClose() {
    ren.DrawRect(0, 0, 500, 500, [3]float32{0.1, 0.5, 0.7})

  }
}
