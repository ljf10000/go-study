package main

import (
	"fmt"
	"testing"
)

const BUFSIZE = 128

type bMessage struct {
	count int
	buf   [BUFSIZE]byte
}

type cMessage struct {
	ch chan *bMessage
}

func client(gch chan *cMessage) {
	ch := make(chan *bMessage, 2)

	for {
		cmsg := &cMessage{
			ch: ch,
		}

		select {
		case gch <- cmsg:
		case bmsg, ok := <-ch:
			if ok && 0 == (bmsg.count%300000) {
				fmt.Println("count =", bmsg.count)
			}
		}
	}
}

func server(gch chan *cMessage) {
	for {
		select {
		case cmsg, ok := <-gch:
			if ok {
				count++
				bmsg := &bMessage{
					count: count,
				}

				cmsg.ch <- bmsg
			}
		}
	}
}

const COUNT = 1 * 1024
const SERVER = 128

var GCH [SERVER]chan *cMessage
var count = 0

func Test1(t *testing.T) {
	/*
		for i := 0; i < SERVER; i++ {
			GCH[i] = make(chan *cMessage, COUNT)
		}

		for i := 0; i < SERVER*COUNT; i++ {
			go client(GCH[i%SERVER])
		}

		for i := 0; i < SERVER-1; i++ {
			go server(GCH[i])
		}

		server(GCH[SERVER-1])
	*/
}
