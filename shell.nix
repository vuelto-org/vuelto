{ pkgs ? import <nixpkgs> {} }:

let
  wasmtest = pkgs.writeShellScriptBin "wasmtest" ''
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
  '';
in
pkgs.mkShell {
  buildInputs = [
    pkgs.xorg.libX11
    pkgs.xorg.libXext
    pkgs.xorg.libXrandr
    pkgs.xorg.libXinerama
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libXxf86vm

    pkgs.pkg-config
    pkgs.gnumake

    pkgs.mesa
    pkgs.libglvnd
    pkgs.alsa-lib

    pkgs.go
    pkgs.gcc

    pkgs.gopls
    pkgs.wasmserve
    wasmtest
  ];
}

