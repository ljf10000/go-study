package array

import (
	"fmt"
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
