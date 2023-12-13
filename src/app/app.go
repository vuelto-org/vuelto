package app


type Application struct {
  Window *Window
  MultipleWindow bool
}

func NewApp() Application {
  return Application{nil, false}
}



