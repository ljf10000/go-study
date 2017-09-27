package main

import (
	. "asdf"
	"bufio"
	ipip "study/ipip/libipip"
	"time"

	"sync"

	goip "github.com/17mon/go-ipip/ipip"

	"os"
)

var prefixs = make([]string, 0, ipip.COUNT)

func init() {
	f, err := os.Open("../iploader/prefix.txt")
	if nil != err {
		panic(err.Error())
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	for r.Scan() {
		prefixs = append(prefixs, r.Text())
	}
}

func perfrom(times, co, cocount int) {
	begin := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		for _, ip := range prefixs {
			goip.Find2(ip)
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
	err := goip.Load("mydata4vipday2.datx")
	if nil != err {
		panic("no datx file")
	}

	perfrom(1, 0, 1)
	perfrom(10, 0, 1)

	for count := 10; count < 100; count += 10 {
		for i := 0; i < count; i++ {
			wg.Add(1)

			go func(co, cocount int) {
				perfrom(1, co, cocount)

				wg.Add(-1)
			}(i, count)
		}
		wg.Wait()
	}
}
