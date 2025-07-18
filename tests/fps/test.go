package main

import (
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w, err := vuelto.NewWindow("hi", 800, 600, false, false)
	if err != nil {
		log.Fatalln(err)
	}

	w.SetFPS(60)

	for !w.Close() {
		w.Refresh()
	}
}
