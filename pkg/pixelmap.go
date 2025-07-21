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
	"vuelto.pp.ua/internal/image"
)

type Pixelmap struct {
	Texture  map[int]map[int][4]int
	Renderer *Renderer2D

	Width  int
	Height int

	vertices []float32
	indices  []uint16
	program  *gl.Program
}

// Loads a new pixelmap and returns a Pixelmap struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewPixelmap() (*Pixelmap, error) {
	r.Window.SetCurrent()

	vertexShader, err := gl.NewShader(gl.VertexShader{
		WebShader:     ushaders.WebVShader,
		DesktopShader: ushaders.DesktopVShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new VertexShader for PixelMap\n %s", err)
	}

	fragmentShader, err := gl.NewShader(gl.FragmentShader{
		WebShader:     ushaders.WebFShader,
		DesktopShader: ushaders.DesktopFShader,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create a new FragmentShader for PixelMap\n %s", err)
	}

	vertexShader.Compile()
	defer vertexShader.Delete()

	fragmentShader.Compile()
	defer fragmentShader.Delete()

	program := gl.NewProgram(*vertexShader, *fragmentShader)
	program.Link()
	program.Use()

	vertices := []float32{
		-1, -1, 0.0, 0.0, 0.0,
		-1, -1 - 1, 0.0, 0.0, 1.0,
		-1 + 1, -1 - 1, 0.0, 1.0, 1.0,
		-1 + 1, -1, 0.0, 1.0, 0.0,
	}

	location, err := program.UniformLocation("uniformColor")
	location.Set(0, 0, 0, 1.0)
	if err != nil {
		return nil, fmt.Errorf("Failed to find uniformColor location for PixelMap\n %s", err)
	}

	location, err = program.UniformLocation("useTexture")
	location.Set(1)
	if err != nil {
		return nil, fmt.Errorf("Failed to find useTexture location for PixelMap\n %s", err)
	}

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}
	r.Window.UnsetCurrent()

	return &Pixelmap{
		Renderer: r,
		Width:    r.Window.Width,
		Height:   r.Window.Height,

		vertices: vertices,
		indices:  indices,
		program:  program,
	}, nil
}

// SetPixel assigns an RGBA color to the (x, y) coordinate.
func (p *Pixelmap) SetPixel(x, y int, color [4]int) {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return
	}

	if p.Texture == nil {
		p.Texture = make(map[int]map[int][4]int)
	}

	if p.Texture[x] == nil {
		p.Texture[x] = make(map[int][4]int)
	}

	p.Texture[x][y] = [4]int{color[0], color[1], color[2], color[3]}
}

// Draw renders the pixelmap on the screen.
func (p *Pixelmap) Draw() {
	p.Renderer.Window.SetCurrent()

	p.program.Use()

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(image.LoadPixelmap(p.Texture, p.Renderer.Window.Width, p.Renderer.Window.Height), gl.NEAREST)
	texture.UnBind()

	buffer := gl.GenBuffers(p.vertices, p.indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(p.program)

	p.program.Use()
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	buffer.Update(p.vertices)

	texture.Bind()
	gl.DrawElements(p.indices)
	texture.UnBind()

	buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
	p.program.UnUse()

	p.Renderer.Window.UnsetCurrent()
}
