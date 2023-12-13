package app

import (
  "runtime"
  "github.com/go-gl/glfw/v3.3/glfw"
  
)

type Application struct {
  Window *Window
  MultipleWindow bool
}

func NewApp() Application {
  runtime.LockOSThread()

  err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()


  return Application{nil, false}
}


