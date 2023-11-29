# Linux
You have to have X11, OpenGL, CMake and a C compiler installed.
## X11
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

# Mac
You have to have Xcode command line tools and CMake inatalled.
## Xcode
```bash
xcode-select --install
```

# Windows
You have to have CMake and a C compiler installed. You can also use WSL, if so follow the Linux instructions.

## Building Vuelto
Clone the repository
```bash
git clone --recurse-submodules https://github.com/dimkauzh/vuelto.git
```
Build the repo with cmake (choose toolchain yourself)
```bash
cd vuelto
mkdir build
cd build
cmake ..
```
