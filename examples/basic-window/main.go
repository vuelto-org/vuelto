package main

import vuelto "github.com/vuelto-org/vuelto/pkg"

func main() {
	w := vuelto.NewWindow("hi", 800, 600, false)

	for !w.Close() {
		w.Refresh()
	}
}
