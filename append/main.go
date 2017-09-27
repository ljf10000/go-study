package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	z := []int{}
	n := []int(nil)

	fmt.Println("append s =", append(s, 10))
	fmt.Println("append z =", append(z, 10))
	fmt.Println("append n =", append(n, 10))
}
