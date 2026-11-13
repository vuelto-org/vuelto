//go:build windows || linux || darwin || freebsd || netbsd || openbsd

/*
 * Copyright (C) 2026 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package opengl

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Shader & Program Management
var CreateShader func(shaderType uint32) uint32
var ShaderSource func(shader uint32, count int32, strings **uint8, length *int32)
var CompileShader func(shader uint32)
var GetShaderiv func(shader uint32, pname uint32, params *int32)
var GetShaderInfoLog func(shader uint32, maxLength int32, length *int32, infoLog *uint8)
var DeleteShader func(shader uint32)
var CreateProgram func() uint32
var AttachShader func(program uint32, shader uint32)
var LinkProgram func(program uint32)
var GetProgramiv func(program uint32, pname uint32, params *int32)
var GetProgramInfoLog func(program uint32, maxLength int32, length *int32, infoLog *uint8)
var UseProgram func(program uint32)
var DeleteProgram func(program uint32)

// Uniforms & Attributes
var GetUniformLocation func(program uint32, name *uint8) int32
var GetUniformiv func(program uint32, location int32, params *int32)
var Uniform1f func(location int32, v0 float32)
var Uniform2f func(location int32, v0, v1 float32)
var Uniform3f func(location int32, v0, v1, v2 float32)
var Uniform4f func(location int32, v0, v1, v2, v3 float32)
var VertexAttribPointerWithOffset func(index uint32, size int32, xtype uint32, normalized bool, stride int32, offset int)
var EnableVertexAttribArray func(index uint32)

// VAO / VBO / EBO
var GenVertexArrays func(n int32, arrays *uint32)
var GenBuffers func(n int32, buffers *uint32)
var BindVertexArray func(array uint32)
var BindBuffer func(target uint32, buffer uint32)
var BufferData func(target uint32, size int, data unsafe.Pointer, usage uint32)
var DeleteVertexArrays func(n int32, arrays *uint32)
var DeleteBuffers func(n int32, buffers *uint32)

// Textures
var GenTextures func(n int32, textures *uint32)
var BindTexture func(target uint32, texture uint32)
var TexImage2D func(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
var TexParameteri func(target uint32, pname uint32, param int32)
var DeleteTextures func(n int32, textures *uint32)

// Rendering
var DrawElements func(mode uint32, count int32, xtype uint32, indices unsafe.Pointer)
var DrawArrays func(mode uint32, first int32, count int32)
var Clear func(mask uint32)
var ClearColor func(red, green, blue, alpha float32)
var Enable func(cap uint32)
var BlendFunc func(sfactor, dfactor uint32)
var Viewport func(x, y, width, height int32)

const (
	TEXTURE_2D           = 0x0DE1
	BLEND                = 0x0BE2
	LINEAR               = 0x2601
	NEAREST              = 0x2600
	ARRAY_BUFFER         = 0x8892
	ELEMENT_ARRAY_BUFFER = 0x8893
	VERTEX_ARRAY         = 0x8074
	VERTEX_SHADER        = 0x8B31
	FRAGMENT_SHADER      = 0x8B30
	COMPILE_STATUS       = 0x8B81
	INFO_LOG_LENGTH      = 0x8B84
	LINK_STATUS          = 0x8B82
	FALSE                = 0
	TRUE                 = 1
	COLOR_BUFFER_BIT     = 0x00004000
	DEPTH_BUFFER_BIT     = 0x00000100
	SRC_ALPHA            = 0x0302
	ONE_MINUS_SRC_ALPHA  = 0x0303
	TEXTURE_MIN_FILTER   = 0x2801
	TEXTURE_MAG_FILTER   = 0x2800
	TEXTURE_WRAP_S       = 0x2802
	TEXTURE_WRAP_T       = 0x2803
	CLAMP_TO_EDGE        = 0x812F
	RGBA                 = 0x1908
	UNSIGNED_BYTE        = 0x1401
	FLOAT                = 0x1406
	UNSIGNED_INT         = 0x1405
	TRIANGLES            = 0x0004
	LINES                = 0x0001
	TRIANGLE_FAN         = 0x0006
	STATIC_DRAW          = 0x88E4
	DYNAMIC_DRAW         = 0x88E8
)

func Str(s string) *uint8 {
	b := []byte(s)
	if len(b) == 0 || b[len(b)-1] != 0 {
		b = append(b, 0)
	}
	return &b[0]
}

func Strs(strs ...string) (**uint8, func()) {
	cstrs := make([]*uint8, len(strs))
	for i, s := range strs {
		cstrs[i] = Str(s)
	}
	return &cstrs[0], func() {}
}

func Uint16ToInt32(in []uint16) []int32 {
	out := make([]int32, len(in))
	for i, v := range in {
		out[i] = int32(v)
	}
	return out
}

func Ptr(v any) unsafe.Pointer {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice && rv.Len() > 0 {
		return unsafe.Pointer(rv.Index(0).UnsafeAddr())
	}
	return nil
}

func getOpenGLLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/System/Library/Frameworks/OpenGL.framework/OpenGL"
	case "linux", "freebsd", "netbsd", "openbsd":
		return "libGL.so.1"
	case "windows":
		return "opengl32.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func Register() error {
	libGL, err := openLibrary(getOpenGLLibrary())
	if err != nil {
		return fmt.Errorf("failed to load OpenGL library: %w", err)
	}

	// Shader & Program Management
	purego.RegisterLibFunc(&CreateShader, libGL, "glCreateShader")
	purego.RegisterLibFunc(&ShaderSource, libGL, "glShaderSource")
	purego.RegisterLibFunc(&CompileShader, libGL, "glCompileShader")
	purego.RegisterLibFunc(&GetShaderiv, libGL, "glGetShaderiv")
	purego.RegisterLibFunc(&GetShaderInfoLog, libGL, "glGetShaderInfoLog")
	purego.RegisterLibFunc(&DeleteShader, libGL, "glDeleteShader")
	purego.RegisterLibFunc(&CreateProgram, libGL, "glCreateProgram")
	purego.RegisterLibFunc(&AttachShader, libGL, "glAttachShader")
	purego.RegisterLibFunc(&LinkProgram, libGL, "glLinkProgram")
	purego.RegisterLibFunc(&GetProgramiv, libGL, "glGetProgramiv")
	purego.RegisterLibFunc(&GetProgramInfoLog, libGL, "glGetProgramInfoLog")
	purego.RegisterLibFunc(&UseProgram, libGL, "glUseProgram")
	purego.RegisterLibFunc(&DeleteProgram, libGL, "glDeleteProgram")

	// Uniforms & Attributes
	purego.RegisterLibFunc(&GetUniformLocation, libGL, "glGetUniformLocation")
	purego.RegisterLibFunc(&GetUniformiv, libGL, "glGetUniformiv")
	purego.RegisterLibFunc(&Uniform1f, libGL, "glUniform1f")
	purego.RegisterLibFunc(&Uniform2f, libGL, "glUniform2f")
	purego.RegisterLibFunc(&Uniform3f, libGL, "glUniform3f")
	purego.RegisterLibFunc(&Uniform4f, libGL, "glUniform4f")
	purego.RegisterLibFunc(&VertexAttribPointerWithOffset, libGL, "glVertexAttribPointer")
	purego.RegisterLibFunc(&EnableVertexAttribArray, libGL, "glEnableVertexAttribArray")

	// VAO / VBO / EBO
	purego.RegisterLibFunc(&GenVertexArrays, libGL, "glGenVertexArrays")
	purego.RegisterLibFunc(&GenBuffers, libGL, "glGenBuffers")
	purego.RegisterLibFunc(&BindVertexArray, libGL, "glBindVertexArray")
	purego.RegisterLibFunc(&BindBuffer, libGL, "glBindBuffer")
	purego.RegisterLibFunc(&BufferData, libGL, "glBufferData")
	purego.RegisterLibFunc(&DeleteVertexArrays, libGL, "glDeleteVertexArrays")
	purego.RegisterLibFunc(&DeleteBuffers, libGL, "glDeleteBuffers")

	// Textures
	purego.RegisterLibFunc(&GenTextures, libGL, "glGenTextures")
	purego.RegisterLibFunc(&BindTexture, libGL, "glBindTexture")
	purego.RegisterLibFunc(&TexImage2D, libGL, "glTexImage2D")
	purego.RegisterLibFunc(&TexParameteri, libGL, "glTexParameteri")
	purego.RegisterLibFunc(&DeleteTextures, libGL, "glDeleteTextures")

	// Rendering
	purego.RegisterLibFunc(&DrawElements, libGL, "glDrawElements")
	purego.RegisterLibFunc(&DrawArrays, libGL, "glDrawArrays")
	purego.RegisterLibFunc(&Clear, libGL, "glClear")
	purego.RegisterLibFunc(&ClearColor, libGL, "glClearColor")
	purego.RegisterLibFunc(&Enable, libGL, "glEnable")
	purego.RegisterLibFunc(&BlendFunc, libGL, "glBlendFunc")
	purego.RegisterLibFunc(&Viewport, libGL, "glViewport")

	return nil
}
