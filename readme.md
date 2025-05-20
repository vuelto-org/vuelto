<!--  markdownlint-disable md033 -->

<div align="center">
  <img width="1400" alt="banner" src="https://github.com/vuelto-org/vuelto/raw/latest/website/logo/banner-dark.png#gh-dark-mode-only">
  <img width="1400" alt="banner" src="https://github.com/vuelto-org/vuelto/raw/latest/website/logo/banner-light.png#gh-light-mode-only">

[![GitHub last commit](https://img.shields.io/github/last-commit/vuelto-org/vuelto?style=for-the-badge)](https://github.com/vuelto-org/vuelto)
[![License](https://img.shields.io/badge/license-VL--Cv1.1-blue?style=for-the-badge)](https://github.com/vuelto-org/license)
[![CI Check](https://img.shields.io/github/actions/workflow/status/vuelto-org/vuelto/builds.yml?style=for-the-badge)](https://github.com/vuelto-org/vuelto/actions/workflows/ci_check.yml)
[![Lines of code](https://www.aschey.tech/tokei/github/vuelto-org/vuelto?style=for-the-badge)](https://github.com/vuelto-org/vuelto)
[![Report card](https://goreportcard.com/badge/vuelto.pp.ua?style=for-the-badge)](https://goreportcard.com/report/vuelto.pp.ua)
[![Powered By](https://img.shields.io/badge/powered_by-GL_3.3-blue?style=for-the-badge)](https://www.opengl.org/Documentation/Specs.html)
[![Made in Ukraine](https://img.shields.io/badge/made_in-ukraine-ffd700.svg?labelColor=0057b7&style=for-the-badge)](https://beua.today)

</div>

Vuelto is an open-source, fast, simple and lightweight game engine, based on Golang and OpenGL. Designed with simplicity and speed in mind, while still staying relatively lightweight. It can be deployed on almost all major platforms, including Linux, MacOS, Windows and Web.

> [!NOTE]
> While Vuelto has a stable version, it's pretty much experimental software. Want something rock solid? Try [Ebitengine](https://github.com/hajimehoshi/ebiten)!

## ✨ Features

- 🌍 Cross Platform
- 🛠️ Open-Source
- 📚 Easy to learn
- 🚀 Fully built using CGo (and some other libraries)

## 📦 Installation

### 📋 Requirements

You need to have the following installed on your system:

- 🖥️ A C compiler (due to CGo)
- 🔧 A Go compiler (Go 1.23 and above)
- 🪟 On Linux, Xorg/Wayland development packages

For an installation guide, [head over here](https://vuelto.pp.ua/install/).

### 🐹 Go package

You can get the latest Go package by running this command:

```sh
go get vuelto.pp.ua@latest
```

## 🖼️ Vuelto example

```go
package main

import (
 vuelto "vuelto.pp.ua/pkg"
)

func main() {
 w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
 ren := w.NewRenderer2D()

 image := ren.LoadImage("test/image.png", 0, 0, 0.5, 0.5)

 for !w.Close() {
  image.Draw()
  w.Refresh()
 }
}
```

## 🖥️ Platform support

Vuelto has built-in deployment support for the following platforms:

| Platform | Status | Version |
| :---- | :---- | :---- |
| Windows | ✅ | v1.0 |
| macOS (Darwin) | ✅ | v1.0 |
| Linux | ✅ | v1.0 |
| Web | ✅ | v1.1 |

## 📖 Docs

You can check out Vuelto's documentation at [Vuelto's website](https://vuelto.pp.ua/docs/).

> [!TIP]
> In case a part documentation is missing something or there is something wrong, use the [GoDoc](https://pkg.go.dev/vuelto.pp.ua) page for API documentation. Use the [examples](https://github.com/vuelto-org/vuelto/tree/stable/examples) directory for usage examples.

### 🛣️ Roadmap

Our roadmap is available on our [website](https://vuelto.pp.ua/roadmap/), under the Roadmap section.

### 🤝 Contributing

We're really thankful for your contributions! Please see our [Contributing Guide](contributing.md) for details.

1. 🍴 Fork the repository
2. 🌟 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 📝 Commit your changes (`git commit -m 'Add some amazing feature'`)
4. 🚀 Push to the branch (`git push origin feature/amazing-feature`)
5. 🔄 Open a Pull Request

## 🛡️ Support & Security

### 🐛 Issues

See the [Issues](https://github.com/vuelto-org/vuelto/issues) page for current bugs and feature requests. In case you find any bug or have a suggestion, please open up an issue or search for any other form of contact to submit a bug report.

#### 🔒 Security Issues

If you find a security vulnerability, please follow the instructions in [security.md](security.md) to safely report it.

### 🔐 License

Vuelto is licensed under the [VL-Cv1.1 License](license.md). Any PRs that will change the license won't be accepted.

### 📖 Guidelines

We’re excited to have you here! To ensure a welcoming and productive environment, we kindly ask you to follow our **Guidelines**. Please take a moment to review the following:

- **[Contributing](contributing.md)**: Learn how to make meaningful contributions to our project.
- **[Code of Conduct](code_of_conduct.md)**: Understand the principles that foster a respectful and inclusive community.

### 🌐 Community and Contact

You can contact us via our Discord community or at our email:

- 🗨️ [Discord server](https://vuelto.pp.ua/discord)
- ✉️ [Email](mailto:dima@vuelto.pp.ua)

### 🙌 Thanks To

A special thanks to:

- **Dimkauzh** for the initial idea and development.
- **ZakaMakesStuff** for the great improvements on top of vuelto.

Also a big shout-out to our homies and partners at [**Sokora**](https://sokora.org), [**Atom**](https://atomlabs.ie) and [**Epic-Bot**](https://bamb.cl/epic-bot/)

Without the help of these people, Vuelto wouldn't be where it is today. Your support has helped make Vuelto even better! 🤝 🙌

---

<h4 align="center">Made with ❤️ by the Vuelto Team </h4>
