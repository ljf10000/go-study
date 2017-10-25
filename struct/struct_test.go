package main

import (
	"fmt"
	"testing"
)

type IShow interface {
	Show()
}

type head struct {
	h1 int
	h2 int
}

func (me *head) Show() {
	fmt.Println("head's h1=", me.h1)
	fmt.Println("head's h2=", me.h2)
}

type body struct {
	head

	b1 int
	b2 int
}

func (me *body) Show() {
	me.head.Show()

	fmt.Println("body's b1=", me.b1)
	fmt.Println("body's b2=", me.b2)
}

type msg struct {
	n int
	s string
}

func Test1(t *testing.T) {
	var b body
	b.h1 = 10
	b.h2 = 11
	b.b1 = 100
	b.b2 = 101

	b.Show()

	msg := &msg{n: 2, s: "sb"}
	fmt.Println(msg)
}
