package vuelto

import (
	"github.com/dimkauzh/vuelto-engine/src"
)

type Application = src.Application
type Window = src.Window
type Image = src.Image
type Renderer2D = src.Renderer2D
type AudioPlayer = src.AudioPlayer
type Animation = src.Animation

func NewApp() src.Application {
	return src.NewApp()
}
