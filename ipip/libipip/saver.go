package libipip

import (
	. "asdf"
	"encoding/gob"
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

type Saver struct {
	loader *Loader

	indexCache [65536][]IpIndex
	symbolMap  map[string]SymbolDesc
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
	if fields[FieldCountry] == "保留地址" {
		return
	}

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

	for i := Field(0); i < FieldEnd; i++ {
		if !i.Fixed() {
			Log.Info("%s max=%d", i, i.Max())
		}
	}
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
