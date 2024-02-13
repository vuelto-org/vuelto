package gl

/*
#include "gl.h"
*/
import "C"

const (
  TEXTURE_2D         = uint(C.GL_TEXTURE_2D)
  TEXTURE_WRAP_S     = uint(C.GL_TEXTURE_WRAP_S)
  TEXTURE_WRAP_T     = uint(C.GL_TEXTURE_WRAP_T)
  TEXTURE_MIN_FILTER = uint(C.GL_TEXTURE_MIN_FILTER)
  TEXTURE_MAG_FILTER = uint(C.GL_TEXTURE_MAG_FILTER)
  CLAMP_TO_EDGE      = uint(C.GL_CLAMP_TO_EDGE)
  LINEAR             = uint(C.GL_LINEAR)

  LINES              = uint(C.GL_LINES)
  QUADS              = uint(C.GL_QUADS)

  RGBA               = uint(C.GL_RGBA)
  UNSIGNED_BYTE      = uint(C.GL_UNSIGNED_BYTE)
)

func Begin(state uint) {
  C.glBegin(C.uint(state))
}

func End() {
  C.glEnd()
}

func Color3f(r, g, b float32) {
  C.glColor3f(C.float(r), C.float(g), C.float(b))
}

func Vertex2f(x, y float32) {
  C.glVertex2f(C.float(x), C.float(y))
}

func ClearColor(r, g, b, a float32) {
  C.glClearColor(C.float(r), C.float(g), C.float(b), C.float(a))
}

func GenTextures(n int32, textures uint) {
  texture := uint32(textures)
  C.glGenTextures(C.int(n), (*C.uint)(&texture))
}

func DeleteTextures(n int, textures uint) {
  texture := uint32(textures)
  C.glDeleteTextures(C.int(n), (*C.uint)(&texture))
}

func BindTexture(target, texture uint) {
  C.glBindTexture(C.uint(target), C.uint(texture))
}

func TexParameteri(target, pname, param uint) {
  C.glTexParameteri(C.uint(target), C.uint(pname), C.int(param))
}

func TexCoord2f(s, t float32) {
  C.glTexCoord2f(C.float(s), C.float(t))
}

func TexImage2D(target, level, internalFormat uint, width, height int, border, format, typ uint, pixels []byte) {
  C.glTexImage2D(C.uint(target), C.int(level), C.int(internalFormat), C.int(width), C.int(height), C.int(border), C.uint(format), C.uint(typ), C.CBytes(pixels))
}