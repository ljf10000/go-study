package main

import (
	. "asdf"
	"unsafe"
)

type tlv struct {
	t int
	l int
	v int
}

func main() {
	b := make([]tlv, 1000, 10000)
	Log.Info("b len=%d, cap=%d", len(b), cap(b))

	for i := 0; i < 10000; i++ {
		b = append(b, tlv{
			t: 1,
			l: 1,
			v: 1,
		})
		Log.Info("b len=%d, cap=%d", len(b), cap(b))
	}

	t := &tlv{}

	Log.Info("t.t offset=%d size=%d", int(unsafe.Offsetof(t.t)), int(unsafe.Sizeof(t.t)))
	Log.Info("t.l offset=%d size=%d", int(unsafe.Offsetof(t.l)), int(unsafe.Sizeof(t.l)))
	Log.Info("t.v offset=%d size=%d", int(unsafe.Offsetof(t.v)), int(unsafe.Sizeof(t.v)))
}
