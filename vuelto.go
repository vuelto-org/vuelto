package vuelto

import (
  "github.com/dimkauzh/vuelto/src"
  
)

type Application = src.Application
type Window = src.Window
type Image = src.Image
type Renderer2D = src.Renderer2D

func NewApp() src.Application {
  return src.NewApp()
}
