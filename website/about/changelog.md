<!-- markdownlint-disable md024 -->
# Vuelto Changelog

## Unreleased (14/11/2026)

### Breaking changes

- Replaced generated/cgo-based OpenGL bindings with runtime-loaded PureGo bindings.
- Replaced direct GLFW usage with the internal PureGo GLFW binding layer.

### Additions

- Added BSD support paths across runtime and CI (FreeBSD, OpenBSD, NetBSD), including new build tags and BSD test/build jobs.
- Added `VUELTO_OPENGL_PATH` and `VUELTO_GFLW_PATH` to override default dynamic library locations.

### Changes

- Added OS-specific OpenGL/GLFW dynamic loader implementations for Unix and Windows.

### Fixes

- Updated build scripts for CGO-disabled builds with PureGo-specific flags.
- Updated Windows CI to install and wire GLFW binaries before running tests/examples.

## Vuelto 1.2.1 (21/07/2025)

### Fixes

- Updated release metadata/version plumbing in `makefile` and `shell.nix`.

## Vuelto 1.2.0 (21/07/2025)

### Breaking changes

- `LoadImage(...)` now requires an `*ImageOptions` argument and changed call sites accordingly.
- Constructors/loaders now return errors in core APIs (window, image, font, shapes, pixelmap, and audio paths).

### Additions

- Added image-processing options (`Blur`, `Contrast`, `Sharpen`, `Invert`) backed by `internal/image/processing`.
- Added RGBA-oriented image pipeline helpers (`RGBA` storage on images and `LoadRGBA` support).
- Added FPS-focused test coverage.

### Changes

- Moved frame pacing/delta tracking to package-level window refresh logic (`GetDeltaTime`, configurable FPS, and framerate-warning controls).
- Reworked PixelMap initialization/draw flow to reuse shader/program state instead of rebuilding setup every draw.

### Fixes

- Replaced many panic/fatal error paths with returned errors across GL, rendering, UI, windowing, and audio internals.
- Updated CI/workflow and dependency set, including bumps for `golang.org/x/image` and `golang.org/x/sys`.

## Vuelto 1.1.2 (25/01/2025)

### Fixes

- Fixed builds for older GPUs.
- Changed to the OpenGL Compatibility profile from the OpenGL Core profile.

## Vuelto 1.1.1 (05/01/2025)

### Fixes

- Fixed web builds not working.

## Vuelto 1.1.0 (05/01/2025)

### Breaking changes

- Updated to OpenGL 3.3 Core, breaking compatibility with unsupported hardware.
- `DrawLine()` now takes arguments in the order `x1 y1 x2 y2`, instead of `x1 x2 y1 y2`.
- `Image` now uses a `Pos` parameter with `Vector2D` type instead of `X, Y` `float32` parameters.
- `Line` now uses `Pos1` and `Pos2` parameters with `Vector2D` type instead of `X1, Y1, X2, Y2`.
- `Rect` now uses a `Pos` parameter with `Vector2D` type instead of `X, Y`.

### Additions

- Support for targeting the web with WASM and WebGL 2.0.
- Event system with keyboard and mouse support.
- `GetDeltaTime()`.
- Framerate management with `GetFPS()` and `SetFPS()`.

### Changes

- The changes above are breaking changes.
- `Vector2D` and `Vector3D` coordinates are now `float32` instead of `float64`.
- Improved performance.

Changes prior to 1.1.0 are not logged. Use GitHub's diff view for older changes.
