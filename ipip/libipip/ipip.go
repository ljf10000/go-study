package libipip

import (
	. "asdf"
	"bufio"
	"encoding/gob"
	"os"
	"strings"
	"unsafe"
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

func bstring(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
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
		saver.add(fields)
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
