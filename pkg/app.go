package pkg

import "runtime"

func Init() {
	runtime.LockOSThread()
}
