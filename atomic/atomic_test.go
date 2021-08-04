package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	COUNT  = 100 * 1000 * 1000
	WORKER = 10
	M      = 1000
)

var record = Record{}

type Record struct {
	writer uint64
}

func (me *Record) add(v uint64) uint64 {
	return atomic.AddUint64(&me.writer, v)
}

func (me *Record) add2(v uint64) uint64 {
	me.writer++

	return me.writer
}

func worker(id int) {
	runtime.LockOSThread()

	for i := 0; i < COUNT; i++ {
		record.add(1)
	}

	wg.Done()
}

func worker2(id int) {
	runtime.LockOSThread()

	for i := 0; i < COUNT; i++ {
		record.add2(1)
	}

	wg.Done()
}

var wg sync.WaitGroup

func Test1(t *testing.T) {
	for i := 0; i < WORKER; i++ {
		wg.Add(1)

		go worker(i)
	}

	old := time.Now().UnixNano()

	wg.Wait()

	now := time.Now().UnixNano()

	used := now - old

	pps := WORKER * COUNT * M / used
	t.Logf("total: %dMpps, worker: %dMpps\n", pps, pps/WORKER)
}

func Test2(t *testing.T) {
	for i := 0; i < WORKER; i++ {
		wg.Add(1)

		go worker2(i)
	}

	old := time.Now().UnixNano()

	wg.Wait()

	now := time.Now().UnixNano()

	used := now - old

	pps := WORKER * COUNT * M / used
	t.Logf("total: %dMpps, worker: %dMpps\n", pps, pps/WORKER)
}
