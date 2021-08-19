package main

import (
	. "asdf"
	"testing"
	"unsafe"
)

type tlv struct {
	t int
	l int
	v int
}

type array struct {
	v []uint32
}

func stack() {
	sta := make([]int, 0, 100)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = append(sta, 1)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 2)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 3)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = sta[: len(sta)-1 : cap(sta)]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[: len(sta)-1 : cap(sta)]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[: len(sta)-1 : cap(sta)]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
}

func stack2() {
	sta := []int{}
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = append(sta, 1)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 2)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 3)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = append(sta, 1)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 2)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = append(sta, 3)
	Log.Info("len=%d cap=%d", len(sta), cap(sta))

	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
	sta = sta[:len(sta)-1]
	Log.Info("len=%d cap=%d", len(sta), cap(sta))
}

func Test1(t *testing.T) {
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

	stack()
	stack2()

	x := []byte(nil)
	y := []byte{}
	z := []byte{0}
	copy(x, y) // empty==>nil
	copy(y, x) // nil==>empty
	copy(z, x) // nil==>bin
	copy(z, y) // empty==>bin

	Log.Info("nil []byte len=%d", len(x))
	Log.Info("empty []byte len=%d", len(y))
}
