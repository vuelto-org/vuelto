package app

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
  ren2d "github.com/dimkauzh/vuelto/src/renderer2d"
)

type Window struct {
  Application Application
  Window *glfw.Window
  Title string
  Width, Height int
}

func (a Application) NewWindow(title string, width, height int, resizable bool) Window {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Error init gl:", err)

	}
	gl.Ortho(0, float64(width), float64(height), 0, -1, 1)

	return Window{a, window, title, width, height}
}

func (w* Window) SetResizable(resizable bool) {
  if resizable {
    w.Window.SetAttrib(glfw.Resizable, glfw.True)
  } else {
    w.Window.SetAttrib(glfw.Resizable, glfw.False)
  }
}

func (w* Window) ShouldClose() bool {
  for !w.Window.ShouldClose() {
		w.Window.SwapBuffers()
		glfw.PollEvents()
    gl.Clear(gl.COLOR_BUFFER_BIT)

    return true
	}
  return false
}


func (w Window) NewRenderer2D() ren2d.Renderer2D {
  return ren2d.Renderer2D{}
}
