package src

import "github.com/go-gl/gl/v2.1/gl"

type Renderer2D struct {}

func (w Window) NewRenderer2D() Renderer2D {
  return Renderer2D{}
}

func (w Renderer2D) Destroy() {
  for _, i := range ImageArray {
    gl.DeleteTextures(1, &i.texture)
  }
}
