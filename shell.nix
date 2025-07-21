{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # X11 libs, needed for glfw
    xorg.libXi
    xorg.libX11
    xorg.libXext
    xorg.libXrandr
    xorg.libXcursor
    xorg.libXxf86vm
    xorg.libXinerama

    # Build tools
    go
    gcc
    gnumake
    pkg-config

    # Graphics apis
    mesa
    libglvnd
    alsa-lib

    # Utilities for development
    pkgs.gopls
    pkgs.wasmserve

    # Used for Vuelto's website
    python312Packages.mkdocs-material
    python312Packages.mkdocs-redirects
    python312Packages.mike

    (pkgs.writeShellScriptBin "wasmtest" ''
      env -i \
        GOARCH="$GOARCH" \
        GOPATH="$GOPATH" \
        GOROOT="$GOROOT" \
        GOCACHE="$GOCACHE" \
        USER="$USER" \
        XDG_CACHE_HOME="$XDG_CACHE_HOME" \
        HOME="$HOME" \
        PATH="$GOPATH/bin:$PATH" \
        VUELTO_DISABLE_BUILD_ERRORS="true" \
        "$(command -v wasmserve)" "$@"
  '')
  ];
}
