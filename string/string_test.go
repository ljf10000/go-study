// string.go project main.go
package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	var s string

	s = `I Like the "the\t \"傻逼\" book"`
	fmt.Printf("string count:%d\n", len(s))
	for k, v := range s {
		fmt.Printf("\t%d:%c\n", k, v)
	}

	b := []byte(s)
	fmt.Printf("byte count:%d\n", len(b))
	for k, v := range b {
		fmt.Printf("\t%d:%c\n", k, v)
	}

	frag := strings.Split(s, " ")
	fmt.Printf("frag count:%d\n", len(frag))
	for _, v := range frag {
		fmt.Printf("\t%s\n", v)
	}

	var s1 string
	var s2 string = ""
	s3 := "s3"
	const s4 = "s4"
	const s5 string = "s4"
	s6 := "s6"

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

	b1 := []byte(s1)
	var b2 []byte

	fmt.Println(b1)
	if nil == b1 {
		fmt.Println("b1 is nil")
	} else {
		fmt.Println("b1 is not nil")
	}

	fmt.Println(b2)
	if nil == b2 {
		fmt.Println("b2 is nil")
	} else {
		fmt.Println("b2 is not nil")
	}

	b = []byte{0, 1, 2}
	fmt.Println("b", string(b))
}
