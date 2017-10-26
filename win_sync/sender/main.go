package main

import (
	"syscall"
	"time"
	"unsafe"

	. "asdf"
)

var kernel = syscall.MustLoadDLL("kernel32.dll")
var CreateEvent = kernel.MustFindProc("CreateEventW")
var PulseEvent = kernel.MustFindProc("PulseEvent")
var handle uintptr

func init() {
	name, _ := syscall.UTF16PtrFromString("sb")
	handle, _, _ = CreateEvent.Call(0, 0, 0, uintptr(unsafe.Pointer(name)))
	if 0 == handle {
		Panic("call CreateEvent fail")
	}
}

func sender() {
	for {
		PulseEvent.Call(handle)
		time.Sleep(1 * time.Nanosecond)
	}
}

func main() {
	sender()
}
