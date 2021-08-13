package main

import (
	. "asdf"
	"os"
	"sync/atomic"
	"time"
)

const SHM_FILE = "/tmp/mplock.shm"
const (
	SHM_SIZE     = 4 * SizeofK
	TIMES_NONE   = 10 * SizeofM
	TIMES_ATOMIC = 1 * SizeofM
	TIMES_WAIT   = 1 * SizeofK
)

type Lock struct {
	value uint32
}

var flock MMap
var lock *Lock
var self uint32
var total uint32

func main() {
	var err error

	if len(os.Args) > 3 {
		total = (uint32)(Atoi(os.Args[2]))
		self = (uint32)(Atoi(os.Args[3]))
	}

	flock, err = MapOpenEx(SHM_FILE, SHM_SIZE, 0)
	if nil != err {
		Log.Error("open %s error: %s", SHM_FILE, err)
		os.Exit(1)
	}

	lock = (*Lock)(SlicePointer(flock))

	switch os.Args[1] {
	case "none":
		lock_none()
	case "atomic":
		lock_atomic()
	case "wait":
		lock_wait()
	case "read":
		lock_read()
	case "reset":
		lock_reset()
	case "begin":
		lock_begin()
	default:
		os.Exit(1)
	}

	flock.Unmap()
}

func wait_begin() uint32 {
	for {
		v := atomic.LoadUint32(&lock.value)
		if v > 0 {
			return v
		}

		time.Sleep(1)
	}
}

func lock_none() {
	begin := wait_begin()

	Log.Info("lock none[%d] begin: %d", self, begin)

	for i := 0; i < TIMES_NONE; i++ {
		lock.value++
	}

	Log.Info("lock none[%d] end: %d", self, lock.value)
}

func lock_atomic() {
	begin := wait_begin()

	Log.Info("lock atomic[%d] begin: %d", self, begin)

	for i := 0; i < TIMES_ATOMIC; i++ {
		atomic.AddUint32(&lock.value, 1)
	}

	Log.Info("lock atomic[%d] end: %d", self, lock.value)
}

func lock_wait() {
	begin := wait_begin()

	Log.Info("lock wait[%d] begin: %d", self, begin)

	for {
		v := atomic.LoadUint32(&lock.value)
		if v > TIMES_WAIT {
			return
		}

		if v%total == self {
			atomic.CompareAndSwapUint32(&lock.value, v, v+1)
		}
	}

	Log.Info("lock wait[%d] end: %d", self, lock.value)
}

func lock_read() {
	Log.Info("read value: %d", lock.value)
}

func lock_reset() {
	atomic.StoreUint32(&lock.value, 0)

	Log.Info("reset value: 0")
}

func lock_begin() {
	atomic.StoreUint32(&lock.value, 1)

	Log.Info("init value: 1")
}
