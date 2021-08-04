package main

import (
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			if i < 5 {
				<-ch
			} else {
				time.Sleep(time.Second)
			}
		}
	}()

	timer := time.NewTimer(100 * time.Millisecond)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case <-timer.C:
			t.Logf("writer write %d timeout", i)
		}
		timer.Reset(100 * time.Millisecond)
	}
	timer.Stop()
}
