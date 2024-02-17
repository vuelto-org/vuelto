package vuelto

import (
	src "github.com/vuelto-org/vuelto/src"
	vuelto "github.com/vuelto-org/vuelto/src"
)

type Application = vuelto.Application
type Window = src.Window
type Image = src.Image
type Renderer2D = src.Renderer2D
type AudioPlayer = src.AudioPlayer

func NewApp() src.Application {
	return src.NewApp()
}
