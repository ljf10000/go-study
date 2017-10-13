package main

import (
	. "asdf"
	"bufio"
	"os"
	ipip "study/ipip/libipip"
	"sync"
	"time"
)

var prefixs = make([]string, 0, ipip.COUNT)

func init() {
	f, err := os.Open("prefix.txt")
	if nil != err {
		panic(err.Error())
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	for r.Scan() {
		prefixs = append(prefixs, r.Text())
	}
}

func check(loader *ipip.Loader) {
	for _, ip := range ipip.IpTests {
		index := loader.FindS(ip)

		Log.Info("found entry@%d", index.IdxEntry())
		Log.Info("%s", loader.EntryDump(index))
	}
}

var searchOnly = false

func perfrom(loader *ipip.Loader, times, co, cocount int) {
	begin := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		for _, ip := range prefixs {
			index := loader.FindS(ip)

			if !searchOnly {
				var s string

				s = loader.EntryField(index, ipip.FieldCountry)
				s = loader.EntryField(index, ipip.FieldProvince)
				s = loader.EntryField(index, ipip.FieldOrganization)
				s = loader.EntryField(index, ipip.FieldNetwork)
				s = loader.EntryField(index, ipip.FieldLat)
				s = loader.EntryField(index, ipip.FieldLng)
				s = loader.EntryField(index, ipip.FieldTimeZone)
				s = loader.EntryField(index, ipip.FieldUTC)
				s = loader.EntryField(index, ipip.FieldRegionalismCode)
				s = loader.EntryField(index, ipip.FieldPhoneCode)
				s = loader.EntryField(index, ipip.FieldCountryCode)
				s = loader.EntryField(index, ipip.FieldContinentCode)
				s = s
			}
		}
	}
	end := time.Now().UnixNano()
	ns := end - begin

	count := len(prefixs)
	Log.Info("%d/%d times:%d, search:%d, all-time:%d ms, one-time:%d ns",
		co,
		cocount,
		times,
		times*count,
		ns/1000000,
		ns/int64(times*count),
	)
}

var wg sync.WaitGroup

func main() {
	loader, err := ipip.Load("mydata4vipday2.sb")
	if nil != err {
		panic(err.Error())
	}
	Log.Info("%s", loader)

	check(loader)
	perfrom(loader, 1, 0, 1)
	perfrom(loader, 10, 0, 1)

	for count := 10; count < 100; count += 10 {
		for i := 0; i < count; i++ {
			wg.Add(1)

			go func(co, cocount int) {
				perfrom(loader, 1, co, cocount)

				wg.Add(-1)
			}(i, count)
		}
		wg.Wait()
	}
}
