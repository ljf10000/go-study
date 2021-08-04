package array

import (
	. "asdf"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type son struct {
	name string
}

type father struct {
	name string
	a    int
	s    son
}

type tlv struct {
	b byte
	s uint16
	i uint32
	I uint
	L uint64

	a [253]byte
}

func TestTlv(t *testing.T) {
	var obj tlv = tlv{}

	fmt.Println(unsafe.Alignof(obj.b))
	fmt.Println(unsafe.Alignof(obj.s))
	fmt.Println(unsafe.Alignof(obj.i))
	fmt.Println(unsafe.Alignof(obj.I))
	fmt.Println(unsafe.Alignof(obj.L))
	fmt.Println(unsafe.Alignof(obj.a))
	fmt.Println("-----")
	fmt.Println(unsafe.Offsetof(obj.b))
	fmt.Println(unsafe.Offsetof(obj.s))
	fmt.Println(unsafe.Offsetof(obj.i))
	fmt.Println(unsafe.Offsetof(obj.I))
	fmt.Println(unsafe.Offsetof(obj.L))
	fmt.Println(unsafe.Offsetof(obj.a))

}

func TestUnsafe(t *testing.T) {
	f := father{
		name: "father",
		s: son{
			name: "son",
		},
	}

	fmt.Println(f)

	fmt.Println(unsafe.Offsetof(f.s))
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(f.s))

	a := [4]father{}
	fmt.Println((uintptr(unsafe.Pointer(&a[3])) - uintptr(unsafe.Pointer(&a[0]))) / unsafe.Sizeof(*(&a[0])))
}

type Proto1 struct {
	a byte
}

type Proto2 struct {
	Proto1
	b byte
}

type Proto3 struct {
	Proto2
	c byte
}

type Proto4 struct {
	Proto3
	d byte
}

type Proto5 struct {
	Proto4

	e uint16
}

type Proto6 struct {
	Proto4

	e uint32
}

func TestStruct(t *testing.T) {
	p1 := Proto1{}
	p2 := Proto2{}
	p3 := Proto3{}
	p4 := Proto4{}
	p5 := Proto5{}
	p6 := Proto6{}

	fmt.Printf("sizeof(p1)=%d\n", unsafe.Sizeof(p1))
	fmt.Printf("sizeof(p2)=%d\n", unsafe.Sizeof(p2))
	fmt.Printf("sizeof(p3)=%d\n", unsafe.Sizeof(p3))
	fmt.Printf("sizeof(p4)=%d\n", unsafe.Sizeof(p4))
	fmt.Printf("sizeof(p5)=%d\n", unsafe.Sizeof(p5))
	fmt.Printf("sizeof(p6)=%d\n", unsafe.Sizeof(p6))
}

type StringHeader struct {
	Data uintptr
	Size int
}

type _Empty struct{}

type _NotAlign struct {
	a uint32
	b byte
	c uint32
	e uint16
	f uint64
}

type small_1 struct {
	a byte
	b byte
	c byte
}

type small_2 struct {
	a uint16
	b byte
}

type small_3 struct {
	a byte
	b uint16
}

func TestString(t *testing.T) {
	s := "0123456789"
	h := (*StringHeader)(unsafe.Pointer(&s))
	b := StructSliceEx(h.Data, h.Size)

	b2 := []byte(s)
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&b2))

	s3 := string(b2)
	h3 := (*StringHeader)(unsafe.Pointer(&s3))

	fmt.Printf("sizeof(s)=%d\n", unsafe.Sizeof(s))
	fmt.Printf("slice(s)=%v\n", b)

	fmt.Printf("h data=0x%x, size=%d\n", h.Data, h.Size)
	fmt.Printf("h2 data=0x%x, size=%d\n", h2.Data, h2.Len)
	fmt.Printf("h3 data=0x%x, size=%d\n", h3.Data, h3.Size)

	empty := _Empty{}
	pEmpty := &_Empty{}
	fmt.Printf("empty size=%d %d %d %d\n",
		unsafe.Sizeof(&empty),
		unsafe.Sizeof(empty),
		unsafe.Sizeof(pEmpty),
		unsafe.Sizeof(*pEmpty))

	var vi interface{}

	fmt.Printf("interface size=%d %d\n",
		unsafe.Sizeof(&vi),
		unsafe.Sizeof(vi))

	bb := true
	fmt.Printf("bool size=%d %d\n",
		unsafe.Sizeof(bb),
		unsafe.Sizeof(false))

	na := _NotAlign{}
	pNA := &na

	fmt.Printf("not-align size=%d %d %d %d\n",
		unsafe.Sizeof(&na),
		unsafe.Sizeof(na),
		unsafe.Sizeof(pNA),
		unsafe.Sizeof(*pNA))

	sm1 := small_1{}
	sm2 := small_2{}
	sm3 := small_3{}

	fmt.Printf("sm1=%d sm2=%d sm3=%d\n",
		unsafe.Sizeof(sm1),
		unsafe.Sizeof(sm2),
		unsafe.Sizeof(sm3))
}
