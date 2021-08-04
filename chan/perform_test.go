package main

import (
	. "asdf"
	"sync"
	"testing"
)

const (
	CHAN_COUNT    = 1024
	MESSAGE_COUNT = 1 * 1000 * 1000

	ENTRY_SIZE   = 104
	ENTRY_COUNT  = 28
	MESSAGE_SIZE = ENTRY_SIZE * ENTRY_COUNT
)

type Entry [ENTRY_SIZE]byte

type Message struct {
	Count uint32
	Entry [ENTRY_COUNT]Entry
}

var wg = &sync.WaitGroup{}

var chBatch = make(chan Message, CHAN_COUNT)

func recvBatch() {
	for {
		msg := <-chBatch

		if 1 == msg.Entry[0][0] {
			break
		}
	}

	wg.Done()
}

func sendBatch() {
	for i := 0; i < MESSAGE_COUNT; i++ {
		msg := Message{}

		chBatch <- msg
	}

	msg := Message{}
	msg.Entry[0][0] = 1

	chBatch <- msg
}

func showBatch(t *testing.T, now, old, ms, us, ns Time64) {
	t.Logf("count:%d us:%d speed:%d/us %d/ms %dK/s",
		MESSAGE_COUNT,
		us,
		ENTRY_COUNT*MESSAGE_COUNT/us,
		ENTRY_COUNT*MESSAGE_COUNT/ms,
		1000*ENTRY_COUNT*MESSAGE_COUNT/us)
}

var chPointer = make(chan *Message, CHAN_COUNT)

func recvPointer() {
	for {
		msg := <-chPointer

		if 1 == msg.Entry[0][0] {
			break
		}
	}

	wg.Done()
}

func sendPointer() {
	for i := 0; i < MESSAGE_COUNT; i++ {
		msg := &Message{}

		chPointer <- msg
	}

	msg := &Message{}
	msg.Entry[0][0] = 1

	chPointer <- msg
}

func showPointer(t *testing.T, now, old, ms, us, ns Time64) {
	t.Logf("count:%d us:%d speed:%d/us %d/ms %dK/s",
		MESSAGE_COUNT,
		us,
		ENTRY_COUNT*MESSAGE_COUNT/us,
		ENTRY_COUNT*MESSAGE_COUNT/ms,
		1000*ENTRY_COUNT*MESSAGE_COUNT/us)
}

var chInt = make(chan int, CHAN_COUNT)

func recvInt() {
	for {
		v := <-chInt

		if -1 == v {
			break
		}
	}

	wg.Done()
}

func sendInt() {
	for i := 0; i < MESSAGE_COUNT; i++ {
		chInt <- i
	}

	chInt <- -1
}

func showInt(t *testing.T, now, old, ms, us, ns Time64) {
	t.Logf("count:%d us:%d speed:%d/us %d/ms %dK/s",
		MESSAGE_COUNT,
		us,
		MESSAGE_COUNT/us,
		MESSAGE_COUNT/ms,
		1000*MESSAGE_COUNT/us)
}

var recvers = []func(){
	0: recvBatch,
	1: recvPointer,
	2: recvInt,
}

var senders = []func(){
	0: sendBatch,
	1: sendPointer,
	2: sendInt,
}

var shows = []func(t *testing.T, now, old, ms, us, ns Time64){
	0: showBatch,
	1: showPointer,
	2: showInt,
}

var IDX = 0

func TestPerform(t *testing.T) {
	wg.Add(1)

	old := NowTime64()

	go recvers[IDX]()
	senders[IDX]()

	wg.Wait()

	now := NowTime64()
	ns := now - old
	us := ns / 1000
	ms := us / 1000

	shows[IDX](t, now, old, ms, us, ns)
}
