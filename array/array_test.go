package array

import (
	"fmt"
	"testing"
)


func TestArray(t *testing.T){
	a := [256]int{}
	fmt.Println("a len =", len(a))
	
	s := a[0:0]
	fmt.Println("s len =", len(s))
	fmt.Println("s cap =", cap(s))
	
	s = append(s, 0)
	fmt.Println("s len =", len(s))
	fmt.Println("s cap =", cap(s))
	
	s = append(s, 1)
	fmt.Println("s len =", len(s))
	fmt.Println("s cap =", cap(s))
}