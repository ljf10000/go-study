package main

import (
	"fmt"
	"testing"
)

func master(v int) (func() int, func(int)) {
	a := v

	return func() int {
			return a
		},
		func(x int) {
			a = x
		}
}

func Test1(t *testing.T) {
	get, set := master(5)

	fmt.Println("get", get())
	set(6)
	fmt.Println("set 6")
	fmt.Println("get", get())
}
