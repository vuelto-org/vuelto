package main

import (
	"github.com/dimkauzh/vuelto-engine"
)

func main() {
	a := vuelto.NewApp()
	w := a.NewWindow("hi", 800, 600, false)

	for !w.Close() {
		w.Refresh()
	}
}
