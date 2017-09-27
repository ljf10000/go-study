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
	UTC             int8         // UTC
	_r              int8         // 保留
	RegionalismCode [6]byte      // 中国行政区划代码
	PhoneCode       [4]byte      // 国际电话代码
	CountryCode     [2]byte      // 国家二位代码
	ContinentCode   [2]byte      // 世界大洲代码 [AF]非洲, [EU]欧洲, [AS]亚洲, [OA]大洋洲, [NA]北美洲, [SA]南美洲, [AN]南极洲
}
*/

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
