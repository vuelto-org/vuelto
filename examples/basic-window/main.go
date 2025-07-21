// This works in the web too! Only it wouldn't be so fun :(
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

	for !w.Close() {
		w.Refresh()
	}
}
