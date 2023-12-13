package vuelto

import (
  "runtime"
  "github.com/go-gl/glfw/v3.3/glfw"
  "github.com/dimkauzh/festi/src/app"
  
)

func NewApp() app.Application {
  runtime.LockOSThread()

  err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()


  return app.Application{nil, false}
}


