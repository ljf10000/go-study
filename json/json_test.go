package main

import (
	"bytes"
	"encoding/gob"
	//	. "asdf"
	. "asdf"
	// "encoding/json"
	"fmt"
	//	. "libhello"
	"testing"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const jperson = `{
	"name":"sb",
	"age":30,
	"live": true,
	"company":[
		{
			"name":"sb company7",
			"age":7,
			"live": true
		},
		{
			"name":"sb company9",
			"age":9,
			"live": true
		}
	],
	"friend":{
		"sb1":"sb1",
		"sb2":"sb2",
		"sb3":"sb3"
	},
	"like":{
		"dsb1":"dsb1",
		"dsb2":"dsb2",
		"dsb3":"dsb3"
	}
}
`
const aperson = `{
	"asdf":[
		{
			"name":"company3",
			"age":3
		},
		{
			"name":"company5",
			"age":5
		}
	]
}
`

type company struct {
	Name string
	Age  int
	Live bool
}

type company2 struct {
	Asdf []company
}

type myJson struct {
	company

	//Company []*company
	//Friend  map[string]string
	//Like    string
	//Error error
	Live bool
}

func Test1(t *testing.T) {
	fmt.Println("jperson", jperson)

	my := &myJson{}

	err := json.Unmarshal([]byte(jperson), my)
	if nil != err {
		fmt.Println("jperson==>my error:", err)
	}
	fmt.Println("my", my)

	my2 := &myJson{
		Live: true,
	}
	fmt.Println("my2", my2)

	ap := &company2{}

	json.Unmarshal([]byte(aperson), ap)
	fmt.Println(ap)

	j, _ := json.Marshal(my)
	fmt.Println(string(j))

	/*
		s := `
			{
				"name":"sb",
				"age":250,
				"alive":true
			}
		`
		Log.Debug("len=%d, s=%s", len(s), s)
		HelloInit(true)
			b := []byte(s)
			b = HelloEncode(b)
			s = string(b)
			Log.Debug("len=%d, s=%s", len(s), s)

			b, _ = HelloDecode(b)
			s = string(b)
			Log.Debug("len=%d, s=%s", len(s), s)
	*/

	companys := []*company{
		&company{
			Name: "sb1",
			Age:  1,
			Live: true,
		},
		&company{
			Name: "sb2",
			Age:  2,
			Live: true,
		},
		&company{
			Name: "sb3",
			Age:  3,
			Live: true,
		},
	}

	var bin bytes.Buffer

	enc := gob.NewEncoder(&bin)
	enc.Encode(companys)
	Log.Info("after gob encode, len:%d cap:%d bin:%v", bin.Len(), bin.Cap(), bin.Bytes())

	var companys2 []company
	dec := gob.NewDecoder(&bin)
	dec.Decode(companys2)
	Log.Info("after gob decode, companys2:%v", companys2)
}

type InfoA struct {
	A string
}

type InfoB struct {
	B string
}

type Info struct {
	InfoA
	InfoB
	C string
}

func Test2(t *testing.T) {
	v1 := &Info{
		InfoA: InfoA{
			A: "A",
		},
		InfoB: InfoB{
			B: "B",
		},
		C: "C",
	}

	b1, _ := json.Marshal(v1)
	fmt.Printf("obj==>json: %s\n", string(b1))

	v2 := &Info{}
	json.Unmarshal(b1, v2)
	fmt.Printf("json==>obj: %x\n", v2)

	b2, _ := json.Marshal(v2)
	fmt.Printf("obj==>json: %s\n", string(b2))

	v3 := map[uint32]string{
		1: "s1",
		2: "s2",
		3: "s3",
	}

	b3, _ := json.Marshal(v3)
	fmt.Printf("map==>json: %s\n", string(b3))

	v4 := map[uint32]string{}

	json.Unmarshal(b3, &v4)
	fmt.Printf("json==>map: %v\n", v4)

	v5 := map[int][]byte{
		1: []byte("111"),
		2: []byte("222"),
		3: []byte("333"),
	}

	b5, _ := json.Marshal(v5)
	fmt.Printf("map==>json: %s\n", string(b5))
}
