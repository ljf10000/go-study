package main

import (
	"fmt"
	"testing"
)

type ETest int

const (
	ETa ETest = 1
	ETb ETest = 2
	ETc ETest = 3
)

func (me ETest) ToString() string {
	return "Enum ETest"
}

func (me ETest) String() string {
	return "ETest String"
}

func Test1(t *testing.T) {
	fmt.Println(ETa)
	fmt.Println(ETb)
	fmt.Println(ETc.ToString())
}
