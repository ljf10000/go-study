package main

import (
	. "asdf"
	"runtime"
	"time"
	// "time"
)

const WORKER_COUNT = 128

type Worker struct {
	idx int

	ch chan IMessage
}

func (me *Worker) main() {
	Log.Info("worker[%d] start ...", me.idx)

	for {
		select {
		case msg := <-me.ch:
			msg.Handle()
		}
	}
}

var workers [WORKER_COUNT]Worker

func gc() {
	ticker := time.NewTicker(100 * time.Microsecond)

	for {
		select {
		case <-ticker.C:
			runtime.GC()
		}
	}
}

func main() {
	go gc()

	for i := 0; i < WORKER_COUNT; i++ {
		worker := &workers[i]
		worker.ch = make(chan IMessage, 1024)
		worker.idx = i

		go worker.main()
	}

	count := uint64(0)

	for {
		for i := 0; i < WORKER_COUNT; i++ {
			count++

			size := SizeofMessage + RandSeed.Uint32()%(64*SizeofK) + 1
			buf := make([]byte, size)
			buf[0] = 1

			msg := (*Message)(SlicePointer(buf))
			msg.counter = count
			msg.total = uint64(size)

			wriper := &MsgWriper{
				// buf:     buf,
				// Message: msg,
				ptr: SlicePointer(buf),
			}

			worker := &workers[i]
			worker.ch <- wriper

			if 0 == count%8192 {
				Log.Info("send msg[%d] to worker[%d]", count, i)
			}
		}
	}
}
