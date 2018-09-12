package main

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	input := "00:11:22:33:44:55"
	mac := [6]byte{}
	sep := byte(0)

	fmt.Sscanf(input, "%d%c%d%c%d%c%d%c%d%c%d",
		&mac[0], &sep,
		&mac[1], &sep,
		&mac[2], &sep,
		&mac[3], &sep,
		&mac[4], &sep,
		&mac[5])

	fmt.Println(mac)

	fmt.Println("AABBCCDDEEFF" + time.Now().Format("20060102150405"))
}

func Test2(t *testing.T) {
	fmt.Printf("%03d\n", 99)
	fmt.Printf("%s\n", fmt.Sprintf("%03d", 98))
}

func Test3(t *testing.T) {
	nid := 0

	fmt.Sscanf("note001", "note%d", &nid)
	fmt.Printf("nid=%d\n", nid)
}

type Identify uint64

func (me Identify) String() string {
	return fmt.Sprintf("%d", me)
}

func (me Identify) Show() {
	fmt.Printf("identify(%d:%x)\n", me, me)
}

func Test4(t *testing.T) {
	v := Identify(99)

	v.Show()
}

func test5_helper(t *testing.T, name string, a ...interface{}) {
	if a[0] == nil {
		fmt.Println(name, "== nil")
	} else {
		fmt.Println(name, "is", a[0])
	}
}

func Test5(t *testing.T) {
	var a chan int
	var b []int
	var c map[int]int
	var d func()
	var e interface{}
	var f = c

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e, "f=", f)

	test5_helper(t, "a", a)
	test5_helper(t, "b", b)
	test5_helper(t, "c", c)
	test5_helper(t, "d", d)
	test5_helper(t, "e", e)
	test5_helper(t, "f", f)
}
