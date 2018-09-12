package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Recv() {
	for {
		v := <-ch1

		if v%1000000 == 0 {
			//fmt.Printf("recv=%d\n", v)
		}

		if v+1 == COUNT {
			break
		}
	}

	wg1.Add(-1)
}

var ch1 = make(chan int, 1000)
var wg1 sync.WaitGroup

const COUNT = 1000 * 1000 * 100

func Test1(t *testing.T) {

	wg1.Add(1)
	go Recv()

	begin := time.Now().UnixNano()
	for i := 0; i < COUNT; i++ {
		ch1 <- i
	}
	wg1.Wait()
	end := time.Now().UnixNano()

	diff := end - begin
	s := diff / int64(time.Second)
	ms := diff / int64(time.Millisecond)
	us := diff / int64(time.Microsecond)
	ns := diff / int64(time.Nanosecond)

	fmt.Printf("second=%d, ms=%d, us=%d, ns=%d\n", s, ms, us, ns)
}
