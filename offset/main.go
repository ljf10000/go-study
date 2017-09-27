// offset project main.go
package main

import (
	"fmt"
	"unsafe"
)

type tlv struct {
	b byte
	s uint16
	i uint32
	I uint
	L uint64
	
	a [253]byte
}

func main() {
	var obj tlv = tlv{};
	
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
