# Linux
On Debian and derivates like Ubuntu and Linux Mint the xorg-dev meta-package pulls in the development packages for all of X11.
```bash
sudo apt install xorg-dev
```

On Fedora and derivatives like Red Hat the X11 extension packages libXcursor-devel, libXi-devel, libXinerama-devel and libXrandr-devel required by GLFW pull in all its other dependencies.
```bash
sudo dnf install libXcursor-devel libXi-devel libXinerama-devel libXrandr-devel
```

On FreeBSD the X11 headers are installed along the end-user X11 packages, so if you have an X server running you should have the headers as well. If not, install the xorgproto package.
```bash
pkg install xorgproto
```

On Cygwin the libXcursor-devel, libXi-devel, libXinerama-devel, libXrandr-devel and libXrender-devel packages in the Libs section of the GUI installer will install all the headers and other development related files GLFW requires for X11.

# MacOS
On macOS, there is no need to install anything

# Windows
There is no support for windows for now, but you can install wsl and follow the [linux](#linux) installation. (Depends on your distro)