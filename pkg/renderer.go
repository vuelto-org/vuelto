package pkg

type Renderer2D struct {
	window Window
}

// Creates a new renderer thats connected to the window.
func (w Window) NewRenderer2D() Renderer2D {
	return Renderer2D{w}
}
