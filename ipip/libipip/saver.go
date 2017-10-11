package libipip

import (
	. "asdf"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

var IpTests = []string{
	"023.031.158.132",
	"050.077.055.164",
	"050.199.176.032",
	"050.199.180.240",
	"050.250.062.096",
	"067.078.093.096",
	"076.046.174.000",
	"108.067.142.252",
	"174.166.240.000",
	"207.172.126.048",
}

func logTestIp(ipt string, iEntry uint32) {
	for _, ip := range IpTests {
		if ip == ipt {
			Log.Info("save %s entry@%d", ip, iEntry)
		}
	}
}

type LL struct {
	Lng string
	Lat string
}

type LLMap map[LL]bool

type LLClass struct {
	all    map[string]LLMap
	single map[string]LLMap
	multi  map[string]LLMap
	export map[string][]LL
}

func newLLClass() *LLClass {
	return &LLClass{
		all:    map[string]LLMap{},
		single: map[string]LLMap{},
		multi:  map[string]LLMap{},
		export: map[string][]LL{},
	}
}

func (me *LLClass) dispatch() {
	for k, v := range me.all {
		if 1 == len(v) {
			me.single[k] = v
		} else {
			me.multi[k] = v
		}

		list := []LL{}
		for k2, _ := range v {
			list = append(list, k2)
		}
		me.export[k] = list
	}
}

func (me *LLClass) dump(tag string) {
	prefix := Empty
	if Empty != tag {
		prefix = tag + " "
	}

	Log.Info(prefix+"single:%d", len(me.single))
	for k, v := range me.single {
		Log.Info(Tab+"%d:%s", len(v), k)
	}

	Log.Info(prefix+"multi:%d", len(me.multi))
	for k, v := range me.multi {
		Log.Info(Tab+"%d:%s", len(v), k)
	}
}

type LLDB struct {
	org    *LLClass
	normal *LLClass
}

func newLLDB() *LLDB {
	return &LLDB{
		org:    newLLClass(),
		normal: newLLClass(),
	}
}

func (me *LLDB) dispatch() {
	me.org.dispatch()
	me.normal.dispatch()
}

func (me *LLDB) dump() {
	me.org.dump("org")
	me.normal.dump(Empty)
}

func (me *LLDB) export(file string) {
	buf, err := json.MarshalIndent(me.normal.export, "", Tab)
	if nil != err {
		Panic("invalid normal export")
	}

	f, err := os.Create(file)
	if nil != err {
		Panic("create %s error:%s", file, err)
	}
	defer f.Close()

	f.WriteString(string(buf))
}

func isOrgString(s string) bool {
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			return true
		}
	}

	return false
}

func (me *LLDB) add(key, lng, lat string) {
	ll := LL{
		Lng: lng,
		Lat: lat,
	}

	class := me.normal
	if isOrgString(key) {
		class = me.org
	}

	if _, ok := class.all[key]; !ok {
		class.all[key] = LLMap{}
	}

	class.all[key][ll] = true
}

type Saver struct {
	loader *Loader

	indexCache [65536][]IpIndex
	symbolMap  map[string]SymbolDesc

	lldb [3]*LLDB
}

func (me *Saver) export() {
	for k, v := range me.lldb {
		v.dispatch()
		v.dump()
		v.export(fmt.Sprintf("level%d.json", k))
	}
}

func (me *Saver) addLL(fields []string) {
	Lng := fields[FieldLng]
	Lat := fields[FieldLat]

	Country := fields[FieldCountry]
	Province := fields[FieldProvince]
	City := fields[FieldCity]

	if "*" != City {
		me.lldb[2].add(Country+"/"+Province+"/"+City, Lng, Lat)
	} else if "*" != Province {
		me.lldb[1].add(Country+"/"+Province, Lng, Lat)
	} else if "*" != Country {
		me.lldb[0].add(Country, Lng, Lat)
	}
}

func newSaver() *Saver {
	return &Saver{
		loader: &Loader{
			Hash:    [65536]IndexDesc{},
			Indexs:  make([]IpIndex, 0, COUNT),
			Entrys:  make([]byte, 0, COUNT*Sizeof_IpEntry),
			Symbols: make([]byte, 0, COUNT),
		},
		indexCache: [65536][]IpIndex{},
		symbolMap:  map[string]SymbolDesc{},
		lldb:       [3]*LLDB{newLLDB(), newLLDB(), newLLDB()},
	}
}

func (me *Saver) saveSymbol(symbol string) SymbolDesc {
	desc, ok := me.symbolMap[symbol]
	if !ok {
		desc = me.loader.pushSymbol(symbol)
		me.symbolMap[symbol] = desc
	}

	return desc
}

func (me *Saver) add(fields []string) {
	// skip it
	switch fields[FieldCountry] {
	case "保留地址",
		"本机地址",
		"本地链路",
		"局域网":
		return
	}

	me.addLL(fields)

	// get ip pair string
	ipstrmin := fields[FieldIpMin]
	ipstrmax := fields[FieldIpMax]

	// get ip pair address, network sort
	ipmin := NewIpB(ipstrmin)
	ipmax := NewIpB(ipstrmax)

	iEntry, entry := me.loader.newEntry()

	setSymbol := func(field Field) {
		s := fields[field]

		desc := me.saveSymbol(s)
		entry.setSymbol(field, desc)
		field.SaveMax(len(s))
	}

	setSymbol(FieldCountry)
	setSymbol(FieldProvince)
	setSymbol(FieldCity)
	setSymbol(FieldOrganization)
	setSymbol(FieldNetwork)
	setSymbol(FieldLng)
	setSymbol(FieldLat)
	setSymbol(FieldTimeZone)
	setSymbol(FieldUTC)

	setSlice := func(field Field) {
		entry.setSlice(field, []byte(fields[field]))
	}

	setSlice(FieldRegionalismCode)
	setSlice(FieldPhoneCode)
	setSlice(FieldCountryCode)
	setSlice(FieldContinentCode)

	minprefix := ipmin.Prefix()
	maxprefix := ipmax.Prefix()

	minsuffix := ipmin.Suffix()
	maxsuffix := ipmax.Suffix()

	if minprefix == maxprefix {
		me.cache(maxprefix, minsuffix, maxsuffix, iEntry)
	} else {
		me.cache(minprefix, minsuffix, 0xffff, iEntry)

		for prefix := minprefix + 1; prefix < maxprefix; prefix++ {
			me.cache(prefix, 0, 0xffff, iEntry)
		}

		me.cache(maxprefix, 0, maxsuffix, iEntry)
	}

	logTestIp(ipstrmin, iEntry)
}

func (me *Saver) cache(prefix, minsuffix, maxsuffix uint16, iEntry uint32) {
	ipindex := MakeIpIndex(minsuffix, maxsuffix, iEntry)

	me.indexCache[prefix] = append(me.indexCache[prefix], ipindex)
}

func (me *Saver) convert() {
	loader := me.loader

	for i := 0; i < 65536; i++ {
		if indexs := me.indexCache[i]; nil != indexs {
			count := len(indexs)
			offset := len(loader.Indexs)
			loader.Indexs = append(loader.Indexs, indexs...)

			loader.Hash[i] = MakeIndexDesc(uint32(offset), uint32(count))
		}
	}

	loader.Indexs = loader.Indexs[:]
	loader.Entrys = loader.Entrys[:]
	loader.Symbols = loader.Symbols[:]

	dumpFieldMax()
}

func (me *Saver) save(file string) error {
	f, err := os.Create(file)
	if nil != err {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)

	Log.Info("%s", me.loader)
	if err := enc.Encode(me.loader); nil != err {
		return err
	}

	return nil
}
