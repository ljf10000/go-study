package libipip

import (
	. "asdf"
	"bufio"
	"encoding/gob"
	"os"
	"strings"
)

/*
http://www.ipip.net/api.html

{
    "ret": "ok",              // ret 值为 ok 时 返回 data 数据 为err时返回msg数据
    "data": [
        "中国",                // 国家
        "天津",                // 省会或直辖市（国内）
        "天津",                // 地区或城市 （国内）
        "",                   // 学校或单位 （国内）
        "鹏博士",              // 运营商字段（只有购买了带有运营商版本的数据库才会有）
        "39.128399",          // 纬度     （每日版本提供）
        "117.185112",         // 经度     （每日版本提供）
        "Asia/Shanghai",      // 时区一, 可能不存在  （每日版本提供）
        "UTC+8",              // 时区二, 可能不存在  （每日版本提供）
        "120000",             // 中国行政区划代码    （每日版本提供）
        "86",                 // 国际电话代码        （每日版本提供）
        "CN",                 // 国家二位代码        （每日版本提供）
        "AP"                  // 世界大洲代码        （每日版本提供）
    ]
}
*/

const (
	COUNT = 1024 * 1024
)

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

type Fields []string

func (me Fields) RegionalismCode() []byte {
	return []byte(me[FieldRegionalismCode])
}

func (me Fields) PhoneCode() []byte {
	return []byte(me[FieldPhoneCode])
}

func (me Fields) CountryCode() []byte {
	return []byte(me[FieldCountryCode])
}

func (me Fields) ContinentCode() []byte {
	return []byte(me[FieldContinentCode])
}

func Convert(src, dst string) {
	f, err := os.Open(src)
	if nil != err {
		panic(err.Error())
	}
	defer f.Close()
	Log.Info("open file:%s", src)

	r := bufio.NewScanner(f)
	saver := newSaver()

	Log.Info("scan file:%s ...", src)
	count := 0
	for r.Scan() {
		line := r.Text()

		fields := strings.Split(line, "\t")
		if FieldEnd != len(fields) {
			Panic("invalid ipip line(%d element)", len(fields))
		}

		count++
		//Log.Info("%d:line:%s", count, line)
		saver.add(Fields(fields))
	}
	Log.Info("scan file:%s ok.", src)

	saver.convert()
	saver.save(dst)
}

func Load(file string) (*Loader, error) {
	f, err := os.Open(file)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	dec := gob.NewDecoder(f)

	var loader Loader
	if err := dec.Decode(&loader); nil != err {
		return nil, err
	} else {
		return &loader, nil
	}
}
