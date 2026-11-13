//go:build windows || linux || darwin || freebsd || netbsd

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

package glfw

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

type Action int
type Key int
type InputMode int

const (
	True                    = 1
	False                   = 0
	Resizable               = 0x00020003
	TransparentFramebuffer  = 0x0002000A
	ContextVersionMajor     = 0x00022002
	ContextVersionMinor     = 0x00022003
	OpenGLProfile           = 0x00022008
	OpenGLCoreProfile       = 0x00032001
	OpenGLCompatProfile     = 0x00032002
	OpenGLForwardCompatible = 0x00022006

	StickyKeysMode InputMode = 0x00030002
	CursorMode     InputMode = 0x00030001

	CursorNormal   = 0x00034001
	CursorHidden   = 0x00034002
	CursorDisabled = 0x00034003

	Press   Action = 1
	Release Action = 0
	Repeat  Action = 2
)

type Window struct {
	handle            uintptr
	resizeCallbackPtr uintptr
}

var (
	cInit                       func() int32
	cTerminate                  func()
	cWindowHint                 func(hint, value int32)
	cCreateWindow               func(width, height int32, title *uint8, monitor, share uintptr) uintptr
	cDestroyWindow              func(window uintptr)
	cSetWindowAttrib            func(window uintptr, attrib, value int32)
	cSetWindowOpacity           func(window uintptr, opacity float32)
	cSetWindowTitle             func(window uintptr, title *uint8)
	cSetWindowSize              func(window uintptr, width, height int32)
	cGetWindowSize              func(window uintptr, width, height *int32)
	cGetFramebufferSize         func(window uintptr, width, height *int32)
	cWindowShouldClose          func(window uintptr) int32
	cPollEvents                 func()
	cSwapBuffers                func(window uintptr)
	cMakeContextCurrent         func(window uintptr)
	cSetFramebufferSizeCallback func(window uintptr, cbfun uintptr) uintptr
	cSetInputMode               func(window uintptr, mode, value int32)
	cGetKey                     func(window uintptr, key int32) int32
	cGetCursorPos               func(window uintptr, xpos, ypos *float64)
	cGetError                   func(description **uint8) int32
)

func Str(str string) *uint8 {
	b := []byte(str)
	if len(b) == 0 || b[len(b)-1] != 0 {
		b = append(b, 0)
	}
	return &b[0]
}

func getGlfwLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "libglfw.3.dylib"
	case "linux", "freebsd", "netbsd", "openbsd":
		return "libglfw.so.3"
	case "windows":
		return "glfw3.dll"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

var initialized bool

func Init() error {
	if initialized {
		return nil
	}

	libGLFW, err := openLibrary(getGlfwLibrary())
	if err != nil {
		return fmt.Errorf("Failed to load GLFW library: %w", err)
	}

	purego.RegisterLibFunc(&cInit, libGLFW, "glfwInit")
	purego.RegisterLibFunc(&cTerminate, libGLFW, "glfwTerminate")
	purego.RegisterLibFunc(&cWindowHint, libGLFW, "glfwWindowHint")
	purego.RegisterLibFunc(&cCreateWindow, libGLFW, "glfwCreateWindow")
	purego.RegisterLibFunc(&cDestroyWindow, libGLFW, "glfwDestroyWindow")
	purego.RegisterLibFunc(&cSetWindowAttrib, libGLFW, "glfwSetWindowAttrib")
	purego.RegisterLibFunc(&cSetWindowOpacity, libGLFW, "glfwSetWindowOpacity")
	purego.RegisterLibFunc(&cSetWindowTitle, libGLFW, "glfwSetWindowTitle")
	purego.RegisterLibFunc(&cSetWindowSize, libGLFW, "glfwSetWindowSize")
	purego.RegisterLibFunc(&cGetWindowSize, libGLFW, "glfwGetWindowSize")
	purego.RegisterLibFunc(&cGetFramebufferSize, libGLFW, "glfwGetFramebufferSize")
	purego.RegisterLibFunc(&cWindowShouldClose, libGLFW, "glfwWindowShouldClose")
	purego.RegisterLibFunc(&cPollEvents, libGLFW, "glfwPollEvents")
	purego.RegisterLibFunc(&cSwapBuffers, libGLFW, "glfwSwapBuffers")
	purego.RegisterLibFunc(&cMakeContextCurrent, libGLFW, "glfwMakeContextCurrent")
	purego.RegisterLibFunc(&cSetFramebufferSizeCallback, libGLFW, "glfwSetFramebufferSizeCallback")
	purego.RegisterLibFunc(&cSetInputMode, libGLFW, "glfwSetInputMode")
	purego.RegisterLibFunc(&cGetKey, libGLFW, "glfwGetKey")
	purego.RegisterLibFunc(&cGetCursorPos, libGLFW, "glfwGetCursorPos")
	purego.RegisterLibFunc(&cGetError, libGLFW, "glfwGetError")

	runtime.LockOSThread()
	if cInit() == False {
		return fmt.Errorf("Failed to initialize GLFW")
	}
	initialized = true
	return nil
}

func Terminate() {
	if cTerminate != nil {
		cTerminate()
	}
}

func GetError() (int32, string) {
	if cGetError == nil {
		return 0, ""
	}

	var desc *byte
	code := cGetError(&desc)
	if code == 0 || desc == nil {
		return code, ""
	}

	length := 0
	for ptr := desc; *ptr != 0; ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1)) {
		length++
	}

	return code, unsafe.String(desc, length)
}

func WindowHint(hint, value int) {
	if cWindowHint != nil {
		cWindowHint(int32(hint), int32(value))
	}
}

func CreateWindow(width, height int, title string, monitor, share uintptr) (*Window, error) {
	if cCreateWindow == nil {
		return nil, fmt.Errorf("GLFW not initialized")
	}

	handle := cCreateWindow(int32(width), int32(height), Str(title), monitor, share)
	if handle == 0 {
		code, msg := GetError()
		return nil, fmt.Errorf("Failed to create GLFW window: code %d (%s)", code, msg)
	}

	return &Window{handle: handle}, nil
}

func PollEvents() {
	if cPollEvents != nil {
		cPollEvents()
	}
}

func DetachCurrentContext() {
	if cMakeContextCurrent != nil {
		cMakeContextCurrent(0)
	}
}

func (w *Window) MakeContextCurrent() {
	if cMakeContextCurrent != nil {
		cMakeContextCurrent(w.handle)
	}
}

func (w *Window) SwapBuffers() {
	if cSwapBuffers != nil {
		cSwapBuffers(w.handle)
	}
}

func (w *Window) ShouldClose() bool {
	if cWindowShouldClose != nil {
		return cWindowShouldClose(w.handle) != False
	}
	return true
}

func (w *Window) Destroy() {
	if cDestroyWindow != nil {
		cDestroyWindow(w.handle)
	}
}

func (w *Window) SetAttrib(attrib, value int) {
	if cSetWindowAttrib != nil {
		cSetWindowAttrib(w.handle, int32(attrib), int32(value))
	}
}

func (w *Window) SetOpacity(opacity float32) {
	if cSetWindowOpacity != nil {
		cSetWindowOpacity(w.handle, opacity)
	}
}

func (w *Window) SetTitle(title string) {
	if cSetWindowTitle != nil {
		cSetWindowTitle(w.handle, Str(title))
	}
}

func (w *Window) SetSize(width, height int) {
	if cSetWindowSize != nil {
		cSetWindowSize(w.handle, int32(width), int32(height))
	}
}

func (w *Window) GetSize() (width, height int) {
	var w32, h32 int32
	if cGetWindowSize != nil {
		cGetWindowSize(w.handle, &w32, &h32)
	}
	return int(w32), int(h32)
}

func (w *Window) SetFramebufferSizeCallback(cbfun func(window *Window, width, height int)) {
	if cSetFramebufferSizeCallback == nil {
		return
	}
	cb := purego.NewCallback(func(win uintptr, width, height int32) {
		cbfun(w, int(width), int(height))
	})
	w.resizeCallbackPtr = cb
	cSetFramebufferSizeCallback(w.handle, cb)
}

func (w *Window) SetInputMode(mode InputMode, value int) {
	if cSetInputMode != nil {
		cSetInputMode(w.handle, int32(mode), int32(value))
	}
}

func (w *Window) GetKey(key Key) Action {
	if cGetKey != nil {
		return Action(cGetKey(w.handle, int32(key)))
	}
	return Release
}

func (w *Window) GetCursorPos() (x, y float64) {
	var xpos, ypos float64
	if cGetCursorPos != nil {
		cGetCursorPos(w.handle, &xpos, &ypos)
	}
	return xpos, ypos
}
