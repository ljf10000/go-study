package main

import (
	. "asdf"
	"testing"
	"time"
)

const (
	SizeofSrc = 1024 * SizeofM
	SizeofDst = 1024 * SizeofM

	TIMES = 10
)

func Test1(t *testing.T) {
	src := [SizeofSrc]byte{}
	dst := [SizeofDst]byte{}

	count := SizeofDst / SizeofSrc

	begin := time.Now().UnixNano()

	for i := 0; i < TIMES; i++ {
		for j := 0; j < count; j++ {
			copy(dst[j*SizeofSrc:], src[:])
		}
	}

	end := time.Now().UnixNano()

	ns := int(end - begin)
	ms := ns / 1000000
	size := SizeofDst * TIMES

	t.Logf("ms=%d ns=%d size=%d speed=%dM/s\n", ms, ns, size, 1000*size/ns)
}
