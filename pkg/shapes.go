/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	"fmt"

	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
)

type Line struct {
	Renderer *Renderer2D
	Pos1     *Vector2D
	Pos2     *Vector2D
	Color    [4]int

	buffer  *gl.Buffer
	program *gl.Program
	indices []uint16
}

type Rect struct {
	Renderer *Renderer2D
	Pos      *Vector2D
	Width    float32
	Height   float32
	Color    [4]int

	buffer  *gl.Buffer
	program *gl.Program
	indices []uint16
}

// Loads a new line and returns a Line struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewLine(x1, y1, x2, y2 float32, color [4]int) (*Line, error) {
	r.Window.SetCurrent()

	vertexShader, err := gl.NewShader(gl.VertexShader{
		WebShader:     ushaders.WebVShader,
		DesktopShader: ushaders.DesktopVShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new VertexShader used for line\n %s", err)
	}

	fragmentShader, err := gl.NewShader(gl.FragmentShader{
		WebShader:     ushaders.WebFShader,
		DesktopShader: ushaders.DesktopFShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new FragmentShader used for line\n %s", err)
	}

	vertexShader.Compile()
	defer vertexShader.Delete()

	fragmentShader.Compile()
	defer fragmentShader.Delete()

	program := gl.NewProgram(*vertexShader, *fragmentShader)
	program.Link()

	program.Use()

	vertices := []float32{
		x1, y1, 0.0,
		x2, y2, 0.0,
	}

	location, err := program.UniformLocation("uniformColor")
	location.Set(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to find uniformColor location for line\n %s", err)
	}

	location, err = program.UniformLocation("useTexture")
	location.Set(0)
	if err != nil {
		return nil, fmt.Errorf("Failed to find useTexture location for line\n %s", err)
	}

	indices := []uint16{
		0, 1,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	r.Window.UnsetCurrent()

	return &Line{
		Renderer: r,
		Pos1:     NewVector2D(x1, y1),
		Pos2:     NewVector2D(x2, y2),
		Color:    color,

		buffer:  buffer,
		program: program,
		indices: indices,
	}, nil
}

// Loads a new rect and returns a Rect struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewRect(x, y, width, height float32, color [4]int) (*Rect, error) {
	r.Window.SetCurrent()

	vertexShader, err := gl.NewShader(gl.VertexShader{
		WebShader:     ushaders.WebVShader,
		DesktopShader: ushaders.DesktopVShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new VertexShader used for rect\n %s", err)
	}

	fragmentShader, err := gl.NewShader(gl.FragmentShader{
		WebShader:     ushaders.WebFShader,
		DesktopShader: ushaders.DesktopFShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new FragmentShader used for rect\n %s", err)
	}

	vertexShader.Compile()
	defer vertexShader.Delete()

	fragmentShader.Compile()
	defer fragmentShader.Delete()

	program := gl.NewProgram(*vertexShader, *fragmentShader)
	program.Link()

	program.Use()

	vertices := []float32{
		x, y, 0.0,
		x, y - height, 0.0,
		x + width, y - height, 0.0,
		x + width, y, 0.0,
	}

	location, err := program.UniformLocation("uniformColor")
	location.Set(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to find uniformColor location used for rect\n %s", err)
	}

	location, err = program.UniformLocation("useTexture")
	location.Set(0)
	if err != nil {
		return nil, fmt.Errorf("Failed to find useTexture location used for rect\n %s", err)
	}

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	r.Window.UnsetCurrent()

	return &Rect{
		Renderer: r,
		Width:    width,
		Height:   height,
		Pos:      NewVector2D(x, y),
		Color:    color,

		buffer:  buffer,
		program: program,
		indices: indices,
	}, nil
}

// Draws the line loaded previously
func (l *Line) Draw() {
	l.Renderer.Window.SetCurrent()

	vertices := []float32{
		l.Pos1.X, l.Pos1.Y, 0.0,
		l.Pos2.X, l.Pos2.Y, 0.0,
	}

	l.program.Use()

	l.buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	l.buffer.Update(vertices)
	gl.DrawElements(l.indices)
	l.buffer.UnBind(gl.VA, gl.VBO, gl.EBO)

	l.program.UnUse()

	l.Renderer.Window.UnsetCurrent()
}

// Draws the rect loaded previously
func (r *Rect) Draw() {
	r.Renderer.Window.SetCurrent()

	vertices := []float32{
		r.Pos.X, r.Pos.Y, 0.0,
		r.Pos.X, r.Pos.Y - r.Height, 0.0,
		r.Pos.X + r.Width, r.Pos.Y - r.Height, 0.0,
		r.Pos.X + r.Width, r.Pos.Y, 0.0,
	}

	r.program.Use()

	r.buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	r.buffer.Update(vertices)
	gl.DrawElements(r.indices)
	r.buffer.UnBind(gl.VA, gl.VBO, gl.EBO)

	r.program.UnUse()

	r.Renderer.Window.UnsetCurrent()
}
