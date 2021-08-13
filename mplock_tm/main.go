package main

import (
	. "asdf"
	"os"
	"sync/atomic"
)

const MPLOCK_TM_SHM = "/tmp/mplock_tm.shm"

const (
	TIMES_STEP = 128 * SizeofK

	TIMES_WRITE = 8 * TIMES_STEP
	TIMES_READ  = 8 * TIMES_STEP

	TIMEOUT = 3
)

type tm_obj_t struct {
	MpLockTm

	v uint32
}

var flock MMap
var tm *tm_obj_t

func (me *tm_obj_t) update(v uint32) {
	me.HandleW(TIMEOUT, func() {
		me.v = v
	})
}

func (me *tm_obj_t) reset() {
	Log.Info("reset value ...")

	me.update(0)

	Log.Info("reset value: 0")
}

func (me *tm_obj_t) init() {
	Log.Info("reset value ...")

	me.update(1)

	Log.Info("reset value: 1")
}

func (me *tm_obj_t) read() {
	Log.Info("read value ...")

	v := uint32(0)

	me.HandleR(TIMEOUT, func() {
		v = me.v
	})

	Log.Info("read value: %v", v)
}

func (me *tm_obj_t) wait() {
	Log.Info("wait ...")

	for {
		v := atomic.LoadUint32(&me.v)
		if v > 0 {
			break
		}

		MpLockPause()
	}

	Log.Info("wait ok.")
}

func (me *tm_obj_t) writer() {
	Log.Info("tm lock writer ...")

	me.wait()

	for i := 0; i < TIMES_WRITE; i++ {
		me.HandleW(TIMEOUT, func() {
			me.v++
		})

		if 0 == (i % TIMES_STEP) {
			Log.Info("tm lock writer step[%d]", i/TIMES_STEP)
		}
	}

	Log.Info("tm lock writer ok.")
}

func (me *tm_obj_t) reader() {
	Log.Info("tm lock reader ...")

	me.wait()

	for i := 0; i < TIMES_READ; i++ {
		me.HandleR(TIMEOUT, func() {
			me.v--
		})

		if 0 == (i % TIMES_STEP) {
			Log.Info("tm lock reader step[%d]", i/TIMES_STEP)
		}
	}

	Log.Info("tm lock reader ok.")
}

func (me *tm_obj_t) read_begin() {
	Log.Info("tm lock read begin ...")

	me.LockR(TIMEOUT)

	v := me.v

	Log.Info("tm lock read[0x%x] begin ok.", v)
}

func (me *tm_obj_t) read_end() {
	Log.Info("tm lock read end ...")

	v := me.v

	me.UnLockR()

	Log.Info("tm lock read[0x%x] end ok.", v)
}

func (me *tm_obj_t) updatebegin() {
	Log.Info("tm lock write begin ...")

	me.LockW(TIMEOUT)
	me.v = 0xffff

	Log.Info("tm lock write begin ok.")
}

func (me *tm_obj_t) updateend() {
	Log.Info("tm lock write end ...")

	me.v = 0
	me.UnLockW()

	Log.Info("tm lock write end ok.")
}

func main() {
	var err error

	act := os.Args[1]

	flock, err = MapOpenEx(MPLOCK_TM_SHM, SizeofPage, 0)
	if nil != err {
		Log.Error("open %s error: %s", MPLOCK_TM_SHM, err)
		os.Exit(1)
	}

	tm = (*tm_obj_t)(SlicePointer(flock))

	// os_dump_buffer(tm, 128);

	switch act {
	case "reset":
		tm.reset()
	case "init":
		tm.init()
	case "read":
		tm.read()
	case "reader":
		tm.reader()
	case "writer":
		tm.writer()
	case "read-begin":
		tm.read_begin()
	case "read-end":
		tm.read_end()
	case "write-begin":
		tm.updatebegin()
	case "write-end":
		tm.updateend()
	default:
		os.Exit(1)
	}
}
