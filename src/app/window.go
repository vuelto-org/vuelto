package app

import "github.com/go-gl/glfw/v3.3/glfw"

type Window struct {
  Window *glfw.Window
  Title string
  Width, Height int
}

func (a* Application) NewWindow(title string, width, height int) Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

  if !a.MultipleWindow {
    window.MakeContextCurrent()
  }

  return Window{window, title, width, height}
}

func (w* Window) WindowShouldClose() bool {
  for !w.Window.ShouldClose() {
		w.Window.SwapBuffers()
		glfw.PollEvents()

    return true
	}
  return false
}
