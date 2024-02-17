package vuelto

import "runtime"

type Application struct {
	Window *Window
}

func NewApp() Application {
	runtime.LockOSThread()
	return Application{nil}
}
