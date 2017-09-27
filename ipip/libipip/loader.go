package libipip

import (
	"fmt"
	"sort"
)

type Loader struct {
	Hash    [65536]IndexDesc
	Indexs  []IpIndex
	Entrys  []byte
	Symbols []byte
}

func (me *Loader) String() string {
	return fmt.Sprintf(
		`loader:
    Symbols size:%d
    Entrys count:%d
    Pairs count:%d
`,
		len(me.Symbols),
		len(me.Entrys)/Sizeof_IpEntry,
		len(me.Indexs))
}

func (me *Loader) EntryField(index IpIndex, field Field) string {
	entry := me.entry(index.IdxEntry())

	if field.Fixed() {
		return string(entry.slice(field))
	} else {
		desc := entry.symbol(field)

		return me.symbol(desc)
	}
}

func (me *Loader) EntryDump(index IpIndex) string {
	return fmt.Sprintf(`Entry:
    Country:%s
    Province:%s
    Organization:%s
    Network:%s
    Lng:%s
    Lat:%s
    TimeZone:%s
    UTC:%s
    RegionalismCode:%s
    PhoneCode:%s
    CountryCode:%s
    ContinentCode:%s
`,
		me.EntryField(index, FieldCountry),
		me.EntryField(index, FieldProvince),
		me.EntryField(index, FieldOrganization),
		me.EntryField(index, FieldNetwork),
		me.EntryField(index, FieldLng),
		me.EntryField(index, FieldLat),
		me.EntryField(index, FieldTimeZone),
		me.EntryField(index, FieldUTC),
		me.EntryField(index, FieldRegionalismCode),
		me.EntryField(index, FieldPhoneCode),
		me.EntryField(index, FieldCountryCode),
		me.EntryField(index, FieldContinentCode))
}

func (me *Loader) FindS(ip string) IpIndex {
	return me.Find(NewIpB(ip))
}

func (me *Loader) Find(ip IpB) IpIndex {
	prefix := ip.Prefix()
	indexDesc := me.Hash[prefix]
	if 0 == indexDesc {
		return 0
	}

	indexs := me.indexs(indexDesc)
	if nil == indexs {
		return 0
	}
	count := len(indexs)

	/*
		for i := 0; i < count; i++ {
			ip2 := MakeIpB(prefix, ip.Suffix())
			ip3 := MakeIpB(prefix, indexs[i].MinSuffix())
			ip4 := MakeIpB(prefix, indexs[i].MaxSuffix())
			Log.Info("%s %s %s %s",
				ip,
				ip2,
				ip3,
				ip4)
		}
	*/

	idx := sort.Search(count, func(i int) bool {
		return ip.Suffix() <= indexs[i].MaxSuffix()
	})

	if idx == count {
		return 0
	}

	index := indexs[idx]

	return index
}

func (me *Loader) indexs(desc IndexDesc) []IpIndex {
	offset := desc.Offset()

	return me.Indexs[offset : offset+desc.Count()]
}

func (me *Loader) entry(iEntry uint32) IpEntry {
	offset := iEntry * Sizeof_IpEntry
	bEntry := me.Entrys[offset : offset+Sizeof_IpEntry]

	return IpEntry(bEntry)
}

func (me *Loader) symbol(desc SymbolDesc) string {
	offset := desc.Offset()
	bSymbol := me.Symbols[offset : offset+desc.Size()]

	return string(bSymbol)
}

func (me *Loader) newEntry() (uint32, IpEntry) {
	bEntry := [Sizeof_IpEntry]byte{}

	iEntry := uint32(len(me.Entrys)) / Sizeof_IpEntry
	me.Entrys = append(me.Entrys, bEntry[:]...)

	return iEntry, me.entry(iEntry)
}

func (me *Loader) pushSymbol(symbol string) SymbolDesc {
	bSymbol := []byte(symbol)

	offset := len(me.Symbols)
	size := len(bSymbol)

	desc := MakeSymbolDesc(uint32(offset), uint32(size))

	me.Symbols = append(me.Symbols, bSymbol...)

	/*
		Log.Info("symbol:%s, offset:%d, size:%d, offset:%d",
			symbol,
			desc.Offset(),
			desc.Size(),
			len(me.Symbols))
	*/

	return desc
}
