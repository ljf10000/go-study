package main

import (
	. "asdf"
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

const (
	KSIZE = 32
	VSIZE = SizeofK
	COUNT = 100 * SizeofK
)

type Object struct {
	DB map[string][]byte
}

var obj = &Object{
	DB: make(map[string][]byte),
}

var cache = &bytes.Buffer{}

func insert() {
	kbuf := [KSIZE]byte{}
	vbuf := [VSIZE]byte{}

	n, err := RandSeed.Read(vbuf[:])
	if nil != err {
		Panic("rand read error: %s", err)
	}
	v := vbuf[:n]

	n, err = RandSeed.Read(kbuf[:])
	if nil != err {
		Panic("rand read error: %s", err)
	}
	k := string(kbuf[:n])

	obj.DB[k] = v
}

func save() {
	before := NowTime64()

	for i := 0; i < COUNT; i++ {
		insert()
	}

	obj.DB["sb"] = []byte("sb")

	enc := gob.NewEncoder(cache)
	err := enc.Encode(obj)
	if nil != err {
		Panic("gob enc error: %s", err)
	}
	after := NowTime64()

	Log.Info("save db count:%dK", len(obj.DB)/SizeofK)
	Log.Info("gob enc len:%dM cap:%dM time:%d ms",
		cache.Len()/SizeofM,
		cache.Cap()/SizeofM,
		after.Diff(before)/1000000)

}

func load() {
	before := NowTime64()

	dec := gob.NewDecoder(cache)
	err := dec.Decode(obj)
	if nil != err {
		Panic("gob dec error: %s", err)
	}
	after := NowTime64()

	Log.Info("load db count:%dK", len(obj.DB)/SizeofK)
	Log.Info("gob dec time:%d ms", after.Diff(before)/1000000)

	Log.Info("load db[sb]=%s", string(obj.DB["sb"]))

	for k, v := range obj.DB {
		Log.Info("k: %s, v: %d", hex.EncodeToString([]byte(k)), len(v))
	}
}

func main() {
	save()
	load()
}
