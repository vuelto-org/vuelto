package src

import "github.com/go-gl/gl/v2.1/gl"

func (r *Renderer2D) DrawRect(x, y, width, height float32, color [3]float32) {
	gl.Begin(gl.QUADS)
	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+width, y)
	gl.Vertex2f(x+width, y+height)
	gl.Vertex2f(x, y+height)
	gl.End()

	gl.Color3f(1.0, 1.0, 1.0)
}

func (r *Renderer2D) ClearColor(color [3]float32) {
	gl.ClearColor(color[0], color[1], color[2], 1)
}

func (r *Renderer2D) DrawLine(x1, x2, y1, y2 float32, color [3]float32) {
	gl.LineWidth(1)
	gl.Begin(gl.LINES)
	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()

	gl.Color3f(1.0, 1.0, 1.0)
}
