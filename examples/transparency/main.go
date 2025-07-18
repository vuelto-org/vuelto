// This does not work in the web, as it doest support transparency :(
package main

import (
	"log"

	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w, err := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false, true) // Change the second argument to true
	if err != nil {
		log.Fatalln(err)
	}

	for !w.Close() {
		w.Refresh()
	}
}
