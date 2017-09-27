package main

import (
	"bytes"
	"encoding/gob"
	//	. "asdf"
	. "asdf"
	"encoding/json"
	"fmt"
	//	. "libhello"
)

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
	//company

	//Company []*company
	//Friend  map[string]string
	//Like    string
	//Error error
	Live bool
}

func main() {
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
