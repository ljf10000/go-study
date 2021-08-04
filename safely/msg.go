package main

import (
	. "asdf"

	"time"
	"unsafe"
)

type IMessage interface {
	Handle()
}

const (
	SizeofMessage = 32
)

type Message struct {
	total   uint64
	_       uint64
	_       uint64
	counter uint64
}

func (me *Message) __body__() uintptr {
	return uintptr(unsafe.Pointer(me)) + SizeofMessage
}

func (me *Message) bodySize() int {
	return int(me.total) - SizeofMessage
}

func (me *Message) Body() []byte {
	return StructSliceEx(me.__body__(), me.bodySize())
}

type MsgWriper struct {
	// *Message
	ptr unsafe.Pointer
}

func (me *MsgWriper) Handle() {
	msg := (*Message)(me.ptr)

	body := msg.Body()
	time.Sleep(time.Duration(body[0]) * time.Millisecond)

	// Log.Info("handle msg[%d] body[%d] ok.", me.counter, len(body))
}
