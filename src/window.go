package src

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
  Application Application
  Window *glfw.Window
  Title string
  Width, Height int
}

func framebuffersizecallback(window *glfw.Window, newWidth, newHeight int) {
  gl.Viewport(0, 0, int32(newWidth), int32(newHeight))
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
  
  window.SetFramebufferSizeCallback(framebuffersizecallback)

  window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Error init gl:", err)

	}

	gl.Ortho(0, float64(width), float64(height), 0, -1, 1)

  gl.Enable(gl.BLEND);
  gl.Enable(gl.TEXTURE_2D);

  gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);

	return Window{a, window, title, width, height}
}

func (w* Window) SetResizable(resizable bool) {
  if resizable {
    w.Window.SetAttrib(glfw.Resizable, glfw.True)
  } else {
    w.Window.SetAttrib(glfw.Resizable, glfw.False)
  }
}

func (w* Window) Close() bool {
  for !w.Window.ShouldClose() {
		glfw.PollEvents()
    return false
	}
  clean()
  return true
}

func (w* Window) Refresh() {
  w.Window.SwapBuffers()
  gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (w* Window) SetContextCurrent() {
  w.Window.MakeContextCurrent()
}

func (w *Window) Destroy() {
  w.Window.Destroy()
}

func clean() {
  for _, i := range ImageArray {
    gl.DeleteTextures(1, &i.texture)
  } 
}

