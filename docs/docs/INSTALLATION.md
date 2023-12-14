---
hide:
  - navigation
---

# Installation

## Linux
You have to have X11-dev-packages and C compiler installed.

### X11
On Debian and derivates like Ubuntu and Linux Mint the xorg-dev meta-package pulls in the development packages for all of X11.
```bash
sudo apt install xorg-dev gcc
```

On Fedora and derivatives like Red Hat the X11 extension packages libXcursor-devel, libXi-devel, libXinerama-devel and libXrandr-devel required by GLFW pull in all its other dependencies.
```bash
sudo dnf install libXcursor-devel libXi-devel libXinerama-devel libXrandr-devel gcc
```

On FreeBSD the X11 headers are installed along the end-user X11 packages, so if you have an X server running you should have the headers as well. If not, install the xorgproto package.
```bash
pkg install xorgproto gcc
```

On Cygwin the libXcursor-devel, libXi-devel, libXinerama-devel, libXrandr-devel and libXrender-devel packages in the Libs section of the GUI installer will install all the headers and other development related files GLFW requires for X11.

## Mac
You have to have Xcode command line tools and CMake inatalled.
## Xcode
```bash
xcode-select --install
```

## Windows
You have to have a C compiler installed. You can also use WSL, if so follow the Linux instructions.

### Go package
We have a Go package published, so run this command to add it to your go.mod:
```bash
go get github.com/dimkauzh/vuelto

```
