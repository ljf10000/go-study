package main

import (
	. "asdf"
	"os"
	"sync/atomic"
)

const MPLOCK_RW_SHM = "/tmp/mplock_rw.shm"

const (
	TIMES_STEP = 128 * SizeofK

	TIMES_WRITE = 8 * TIMES_STEP
	TIMES_READ  = 8 * TIMES_STEP
)

type rw_obj_t struct {
	MpLockRw

	v uint32
}

var flock MMap
var rw *rw_obj_t

func (me *rw_obj_t) update(v uint32) {
	me.HandleW(func() {
		me.v = v
	})
}

func (me *rw_obj_t) reset() {
	Log.Info("reset value ...")

	me.update(0)

	Log.Info("reset value: 0")
}

func (me *rw_obj_t) init() {
	Log.Info("reset value ...")

	me.update(1)

	Log.Info("reset value: 1")
}

func (me *rw_obj_t) read() {
	Log.Info("read value ...")

	v := uint32(0)

	me.HandleR(func() {
		v = me.v
	})

	Log.Info("read value: %v", v)
}

func (me *rw_obj_t) wait() {
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

func (me *rw_obj_t) writer() {
	Log.Info("rw lock writer ...")

	me.wait()

	for i := 0; i < TIMES_WRITE; i++ {
		me.HandleW(func() {
			me.v++
		})

		if 0 == (i % TIMES_STEP) {
			Log.Info("rw lock writer step[%d]", i/TIMES_STEP)
		}
	}

	Log.Info("rw lock writer ok.")
}

func (me *rw_obj_t) reader() {
	Log.Info("rw lock reader ...")

	me.wait()

	for i := 0; i < TIMES_READ; i++ {
		me.HandleR(func() {
			me.v--
		})

		if 0 == (i % TIMES_STEP) {
			Log.Info("rw lock reader step[%d]", i/TIMES_STEP)
		}
	}

	Log.Info("rw lock reader ok.")
}

func (me *rw_obj_t) read_begin() {
	Log.Info("rw lock read begin ...")

	me.LockR()

	v := me.v

	Log.Info("rw lock read[0x%x] begin ok.", v)
}

func (me *rw_obj_t) read_end() {
	Log.Info("rw lock read end ...")

	v := me.v

	me.UnLockR()

	Log.Info("rw lock read[0x%x] end ok.", v)
}

func (me *rw_obj_t) updatebegin() {
	Log.Info("rw lock write begin ...")

	me.LockW()
	me.v = 0xffff

	Log.Info("rw lock write begin ok.")
}

func (me *rw_obj_t) updateend() {
	Log.Info("rw lock write end ...")

	me.v = 0
	me.UnLockW()

	Log.Info("rw lock write end ok.")
}

func main() {
	var err error

	act := os.Args[1]

	flock, err = MapOpenEx(MPLOCK_RW_SHM, SizeofPage, 0)
	if nil != err {
		Log.Error("open %s error: %s", MPLOCK_RW_SHM, err)
		os.Exit(1)
	}

	rw = (*rw_obj_t)(SlicePointer(flock))

	// os_dump_buffer(rw, 128);

	switch act {
	case "reset":
		rw.reset()
	case "init":
		rw.init()
	case "read":
		rw.read()
	case "reader":
		rw.reader()
	case "writer":
		rw.writer()
	case "read-begin":
		rw.read_begin()
	case "read-end":
		rw.read_end()
	case "write-begin":
		rw.updatebegin()
	case "write-end":
		rw.updateend()
	default:
		os.Exit(1)
	}
}
