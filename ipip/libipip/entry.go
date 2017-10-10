package libipip

import (
	"encoding/binary"
)

/*
type IpEntry struct {
	IpMin           SymbolDesc
	IpMax           SymbolDesc
	Country         SymbolDesc // 国家
	Province        SymbolDesc // 省会或直辖市（国内）
	City            SymbolDesc // 地区或城市 （国内）
	Organization    SymbolDesc // 学校或单位 （国内）
	Network         SymbolDesc // 运营商
	Lng             SymbolDesc // 经度
	Lat             SymbolDesc // 维度
	TimeZone        SymbolDesc // 时区
	UTC             SymbolDesc // UTC
	RegionalismCode [6]byte      // 中国行政区划代码
	PhoneCode       [4]byte      // 国际电话代码
	CountryCode     [2]byte      // 国家二位代码
	ContinentCode   [2]byte      // 世界大洲代码 [AF]非洲, [EU]欧洲, [AS]亚洲, [OA]大洋洲, [NA]北美洲, [SA]南美洲, [AN]南极洲
}
*/

const (
	FieldIpMin           = 0
	FieldIpMax           = 1
	FieldCountry         = 2
	FieldProvince        = 3
	FieldCity            = 4
	FieldOrganization    = 5
	FieldNetwork         = 6
	FieldLng             = 7
	FieldLat             = 8
	FieldTimeZone        = 9
	FieldUTC             = 10
	FieldRegionalismCode = 11
	FieldPhoneCode       = 12
	FieldCountryCode     = 13
	FieldContinentCode   = 14
	FieldEnd             = 15
)

type fieldInfo struct {
	offset int
	size   int
	max    int
	name   string
	fixed  bool
}

var fieldInfos = [FieldEnd]fieldInfo{
	FieldIpMin: fieldInfo{
		name: "IpMin",
	},
	FieldIpMax: fieldInfo{
		name: "IpMax",
	},
	FieldCountry: fieldInfo{
		name:   "Country",
		offset: 0,
		size:   4,
	},
	FieldProvince: fieldInfo{
		name:   "Province",
		offset: 4,
		size:   4,
	},
	FieldCity: fieldInfo{
		name:   "City",
		offset: 8,
		size:   4,
	},
	FieldOrganization: fieldInfo{
		name:   "Organization",
		offset: 12,
		size:   4,
	},
	FieldNetwork: fieldInfo{
		name:   "Network",
		offset: 16,
		size:   4,
	},
	FieldLng: fieldInfo{
		name:   "Lng",
		offset: 20,
		size:   4,
	},
	FieldLat: fieldInfo{
		name:   "Lat",
		offset: 24,
		size:   4,
	},
	FieldTimeZone: fieldInfo{
		name:   "TimeZone",
		offset: 28,
		size:   4,
	},
	FieldUTC: fieldInfo{
		name:   "UTC",
		offset: 32,
		size:   4,
	},
	FieldRegionalismCode: fieldInfo{
		name:   "RegionalismCode",
		offset: 36,
		size:   6,
		fixed:  true,
	},
	FieldPhoneCode: fieldInfo{
		name:   "PhoneCode",
		offset: 42,
		size:   4,
		fixed:  true,
	},
	FieldCountryCode: fieldInfo{
		name:   "CountryCode",
		offset: 46,
		size:   2,
		fixed:  true,
	},
	FieldContinentCode: fieldInfo{
		name:   "ContinentCode",
		offset: 48,
		size:   2,
		fixed:  true,
	},
}

type Field int

func (me Field) String() string {
	return fieldInfos[me].name
}

func (me Field) Offset() int {
	return fieldInfos[me].offset
}

func (me Field) Size() int {
	return fieldInfos[me].size
}

func (me Field) Fixed() bool {
	return fieldInfos[me].fixed
}

func (me Field) Max() int {
	return fieldInfos[me].max
}

func (me Field) SaveMax(max int) {
	info := &fieldInfos[me]

	if max > info.max {
		info.max = max
	}
}

const Sizeof_IpEntry = 50

type IpEntry []byte

func (me IpEntry) slice(field Field) []byte {
	offset := field.Offset()

	return me[offset : offset+field.Size()]
}

func (me IpEntry) setSlice(field Field, v []byte) {
	copy(me.slice(field), v)
}

func (me IpEntry) symbol(field Field) SymbolDesc {
	desc := binary.LittleEndian.Uint32(me.slice(field))

	return SymbolDesc(desc)
}

func (me IpEntry) setSymbol(field Field, desc SymbolDesc) {
	binary.LittleEndian.PutUint32(me.slice(field), uint32(desc))
}
