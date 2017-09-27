package array

import (
	"fmt"
	. "unsafe"
	"testing"
)

type son struct {
	name string
}

type father struct {
	name string
	a int
	s son
}

func TestUnsafe(t *testing.T){
	f := father {
		name:"father",
		s:son{
			name:"son",
		},
	}
	
	fmt.Println(f)
	
	fmt.Println(Offsetof(f.s))
	fmt.Println(Sizeof(f))
	fmt.Println(Sizeof(f.s))
	
	a := [4]father{}
	fmt.Println((uintptr(Pointer(&a[3])) - uintptr(Pointer(&a[0])))/Sizeof(*(&a[0])))
}