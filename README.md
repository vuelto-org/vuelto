<h1 align="center">
<p align="center">
<img width="1400" alt="banner" src="https://github.com/vuelto-org/vuelto/assets/106883655/2363d776-2669-41f2-b31f-a235de8bea82">

<h1 align="center">A Go Game Engine build with GLFW and OpenGL.</h1>

<p align="center">
  <a href="https://github.com/vuelto-org/vuelto"><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/vuelto-org/vuelto"></a>
  <a href="https://github.com/vuelto-org/vuelto"><img alt="GitHub license" src="https://img.shields.io/github/license/vuelto-org/vuelto"></a>
  <a href="https://github.com/vuelto-org/vuelto"><img alt="CI Check" src="https://github.com/vuelto-org/vuelto/actions/workflows/ci_check.yml/badge.svg"></a>
  <a href="https://github.com/vuelto-org/vuelto"><img alt="Lines of code" src="https://www.aschey.tech/tokei/github/vuelto-org/vuelto"></a>
  <a href="https://goreportcard.com/report/github.com/vuelto-org/vuelto"><img alt="Report card" src="https://goreportcard.com/badge/github.com/vuelto-org/vuelto"></a>
</p>

</h1>

Vuelto is a Game engine that handles windowing, input, rendering and audio. It build to be lightweight but still powerfull enough to handles hard tasks.

## Table of Contents
 - [Installation](https://vuelto-org.github.io/vuelto/docs/INSTALLATION/)
 - [Usage](#usage)
 - [Contributing](#contributing)
 - [Discord Server](https//discord.gg/gZqdRXbbqg)
 - [License](#license)
 - [About](#about)

## Usage
### Requirements
There are some extra things you will need to use Vuelto, [see them here](https://vuelto-org.github.io/vuelto/docs/INSTALLATION/)

### Go package
We have a Go package published, so run this command to add it to your go.mod:
```bash
go get github.com/vuelto-org/vuelto@latest
```

### Examples
All of our examples are inside the examples directory, so take a look there is you want a example. Here one small example of how easy Vuelto is:
```go
package main

import (
	"github.com/vuelto-org/vuelto"
)

func main() {
	a := vuelto.NewApp()
	w := a.NewWindow("hi", 800, 600, false)
	ren := a.NewRenderer2D()

	image := ren.LoadImage("test/image.png", 300, 300, 250, 250)

	for !w.Close() {
		image.Draw()
		w.Refresh()
	}
}
```

## Discord server
We have a [discord server at this link](https://discord.gg/gZqdRXbbqg)

## Contributing
We are fully open to contributions, but I will check and test your code before merging it into the dev branch. All your code thats accepted will only be merged into the dev branch, and will be later released with the next release.

## License
Vuelto is licensed under the [GPLv3 Licence](LICENSE).

## About
Vuelto is a game engine built with GLFW and OpenGL using Go. My goal is to achieve a good and easy to use game engine for everyone. I am currently working on the engine, and I am not planning to releasing Vuelto (its not ready). I am planning to release a website for the engine that contains a documentation for the engine and a roadmap/changelog. I am also planning to release a couple of tutorials on how to use vuelto. I hope you will enjoy using vuelto. Have fun!
