package src

type Application struct {
	Window *Window
}

func NewApp() Application {
	return Application{nil}
}
