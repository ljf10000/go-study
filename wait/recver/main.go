package main

import (
	"syscall"
	"time"
	"unsafe"

	. "asdf"
)

var kernel = syscall.MustLoadDLL("kernel32.dll")
var OpenEvent = kernel.MustFindProc("OpenEventW")
var handle uintptr

const (
	//COUNT = 10 * 1000 * 1000
	COUNT = 10 * 1000
	//COUNT = 1000
)

func init() {
	name, _ := syscall.UTF16PtrFromString("sb")
	handle, _, _ = OpenEvent.Call(0x00100000, 0, uintptr(unsafe.Pointer(name)))
	if 0 == handle {
		Panic("call OpenEvent fail")
	}
}

func recver() {
	count := 0
	begin := time.Now().UnixNano()
	for i := 0; i < COUNT; i++ {
		//Log.Info("recver %d:WaitForSingleObject ...", i)
		syscall.WaitForSingleObject(syscall.Handle(handle), syscall.INFINITE)
		//Log.Info("recver %d:WaitForSingleObject ok.", i)
		count++
	}
	end := time.Now().UnixNano()
	ns := end - begin

	Log.Info("recver times:%dK all:%dms one:%dns", COUNT/1000, ns/1000000, ns/COUNT)
}

func main() {
	recver()
}
