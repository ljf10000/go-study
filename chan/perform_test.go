package main

import (
	. "asdf"
	"sync"
	"testing"
)

const (
	CHAN_COUNT    = 1024
	MESSAGE_COUNT = 1 * 1000 * 1000

	ENTRY_SIZE   = 96
	ENTRY_COUNT  = 48
	MESSAGE_SIZE = ENTRY_SIZE * ENTRY_COUNT
)

type Entry [ENTRY_SIZE]byte

type Message struct {
	Count uint32
	Entry [ENTRY_COUNT]Entry
}

var ch = make(chan Message, CHAN_COUNT)

var wg = &sync.WaitGroup{}

func recver() {
	for {
		msg := <-ch

		if 1 == msg.Entry[0][0] {
			break
		}
	}

	wg.Done()
}

func sender() {
	for i := 0; i < MESSAGE_COUNT; i++ {
		msg := Message{}

		ch <- msg
	}

	msg := Message{}
	msg.Entry[0][0] = 1

	ch <- msg
}

func TestPerform(t *testing.T) {
	wg.Add(1)

	old := NowTime64()

	go recver()
	sender()

	wg.Wait()

	now := NowTime64()
	nano := now - old
	us := nano / 1000
	ms := us / 1000

	t.Logf("count:%d us:%d speed:%d/us %d/ms %dK/s",
		MESSAGE_COUNT,
		us,
		ENTRY_COUNT*MESSAGE_COUNT/us,
		ENTRY_COUNT*MESSAGE_COUNT/ms,
		1000*ENTRY_COUNT*MESSAGE_COUNT/us)
}
