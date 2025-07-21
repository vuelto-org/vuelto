# 🚀 Get Started

## 🛠️ Requirements

There are some extra things you will need to use Vuelto.

- 🖥️ A C compiler
- 🔧 A Go compiler (Go 1.23 and above)
- 🪟 Xorg/Wayland development packages (For Linux only)
- 🖱️ Supported platform
For a installation guide, [go here](install.md).

## 📦 Go package

We have a Go package published, so run this command to add it to your go.mod:

```bash
go get vuelto.pp.ua@latest
```

## 🌟 Examples

All of our examples are inside the examples directory, so take a look there is you want a example. Here one small example of how easy Vuelto is:

```go
package main

import (
 vuelto "vuelto.pp.ua/pkg"
)

func main() {
 w, _ := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
 ren := w.NewRenderer2D()

 image, _ := ren.LoadImage("test/image.png", 0, 0, 0.5, 0.5)

 for !w.Close() {
  image.Draw()
  w.Refresh()
 }
}
```
