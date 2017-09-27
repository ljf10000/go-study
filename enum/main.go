package main

import (
	"fmt"
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

func main() {
	fmt.Println(ETa)
	fmt.Println(ETb)
	fmt.Println(ETc.ToString())
}
