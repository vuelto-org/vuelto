package pkg

type Renderer2D struct{}

func (w *Window) NewRenderer2D() Renderer2D {
	return Renderer2D{}
}
