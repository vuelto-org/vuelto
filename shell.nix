{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # X11 libs, needed for glfw
    libxi
    libX11
    libXext
    libXrandr
    libXcursor
    libXxf86vm
    libXinerama

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
    python3Packages.mkdocs-material
    python3Packages.mkdocs-redirects
    python3Packages.mike

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
