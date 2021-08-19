package main

import (
	. "asdf"
	"os"
	"strings"
)

const MAX_SIZE = 1024

var mpa = MpArray{Trace: true}
var config = []MpArrayConfig{
	MpArrayConfig{Size: 4, Count: 128},
	MpArrayConfig{Size: 8, Count: 128},
	MpArrayConfig{Size: 16, Count: 128},
	MpArrayConfig{Size: 32, Count: 128},
	MpArrayConfig{Size: 64, Count: 128},
	MpArrayConfig{Size: 128, Count: 128},
	MpArrayConfig{Size: 256, Count: 128},
	MpArrayConfig{Size: 512, Count: 128},
	MpArrayConfig{Size: 1024, Count: 128},
}

type rwbuffer struct {
	size int
	body [MAX_SIZE]byte
}

func (me *rwbuffer) Slice() []byte {
	return me.body[:me.size]
}

func write_all() {
	writer := rwbuffer{}

	for i := 0; i < MAX_SIZE; i++ {
		writer.body[i] = 0xff
	}

	for i := 0; i < len(config); i++ {
		for j := uint32(0); j < config[i].Count; j++ {
			idx := MpArrayIndex{Array: uint32(i), Entry: j}

			writer.size = int(config[i].Size)
			if err := mpa.Write(idx, &writer, 3); nil != err {
				Panic("write all error:%v", err)
			}
		}
	}
}

func read_all() {
	reader := rwbuffer{}

	for i := 0; i < len(config); i++ {
		for j := uint32(0); j < config[i].Count; j++ {
			idx := MpArrayIndex{Array: uint32(i), Entry: j}

			reader.size = int(config[i].Size)
			ok, err := mpa.Read(idx, &reader, 3)
			if nil != err {
				Panic("read all error:%v", err)
			} else if false == ok {
				Panic("read all nothing")
			}
		}
	}
}

func main() {
	both := false

	mpa.Init()

	Log.SetLevel(LogLevelDebug)
	if err := mpa.Open(config); nil != err {
		Log.Debug("mparray open error")
		return
	}

	if 2 == len(os.Args) {
		act := os.Args[1]

		if strings.Contains(act, "w") {
			write_all()
		}

		if strings.Contains(act, "r") {
			read_all()
		}
	} else {
		both = true
	}

	if both {
		write_all()
		read_all()
	}

	mpa.Close()
}
