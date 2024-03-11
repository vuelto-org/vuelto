# 🚀 Get Started
## 🛠️ Requirements
There are some extra things you will need to use Vuelto.

- A C compiler
- A Go compiler
- Xorg development packages (For Linux only)

For a installation guide, [go here](install.md).

## 📦 Go package
We have a Go package published, so run this command to add it to your go.mod:
```bash
go get github.com/vuelto-org/vuelto@latest
```

## 🌟 Examples
All of our examples are inside the examples directory, so take a look there is you want a example. Here one small example of how easy Vuelto is:
```go
package main

import (
	vuelto "github.com/vuelto-org/vuelto/pkg"
)

func main() {
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
	ren := w.NewRenderer2D()

	image := ren.LoadImage("test/image.png", 300, 300, 250, 250)
	image1 := ren.LoadImage("test/image.png", 100, 100, 150, 150)

	for !w.Close() {
		image.Draw()
		image1.Draw()
		w.Refresh()

	}
}

```