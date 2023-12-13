package main

import "github.com/dimkauzh/vuelto"

func main() {
  a := vuelto.NewApp()
  w := a.NewWindow("hi", 800, 600, 1)

  for w.WindowShouldClose() {

  }
}
