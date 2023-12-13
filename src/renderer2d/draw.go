package renderer2d

import "github.com/go-gl/gl/v2.1/gl"

func (r* Renderer2D) DrawRect(x, y, width, height float32, color [3]float32) {
  gl.Begin(gl.QUADS);
  gl.Color3f(color[0], color[1], color[2]);
  gl.Vertex2f(x, y);
  gl.Vertex2f(x + width, y);
  gl.Vertex2f(x + width, y + height);
  gl.Vertex2f(x, y + height);
  gl.End();

  gl.Color3f(1.0, 1.0, 1.0);
  
}
