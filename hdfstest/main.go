package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	//	"sync"
	"syscall"
	"time"

	"github.com/colinmarc/hdfs"
	// leveldb "github.com/syndtr/goleveldb/leveldb"
)

const (
	K = 1024
	M = 1024 * K
	G = 1024 * M

	FACTOR  = 1
	MAXSIZE = G
)

var sts []*hdfsst
var cache [MAXSIZE]byte

//var wg sync.WaitGroup
var begin time.Time
var namenode string
var username string
var client *hdfs.Client

// var db *leveldb.DB

type fraging struct {
	size  int
	count int
	times int
}

var frags = []*fraging{
	/*
		&fraging{
			size:  K,
			count: FACTOR * 4 * M,
			times: 64 * K,
		},
		&fraging{
			size:  K * 4,
			count: FACTOR * M / 4,
			times: 16 * K,
		},
	*/
	&fraging{
		size:  K * 16,
		count: 10, //FACTOR * M / 16,
		times: 4 * K,
	},
	/*
		&fraging{
			size:  K * 64,
			count: FACTOR * M / 64,
			times: K,
		},
		&fraging{
			size:  K * 256,
			count: FACTOR * M / 256,
			times: 256,
		},
		&fraging{
			size:  M,
			count: FACTOR * K,
			times: 64,
		},
		&fraging{
			size:  M * 4,
			count: FACTOR * K / 4,
			times: 16,
		},
		&fraging{
			size:  M * 16,
			count: FACTOR * K / 16,
			times: 4,
		},
	*/
	/*
		&fraging{
			size:  M * 64,
			count: K / 64,
			times: 1,
		},
		&fraging{
			size:  M * 256,
			count: 4 * K / 256,
			times: 1,
		},
		&fraging{
			size:  G,
			count: 4,
			times: 1,
		},
	*/
}

type hdfsst struct {
	times int
	bytes int64

	filename string
	r        *hdfs.FileReader
	w        *hdfs.FileWriter
}

func (me *hdfsst) statistics(bytes int) {
	me.times++
	me.bytes += int64(bytes)
}

func (me *hdfsst) close() {
	if nil != me.r {
		me.r.Close()
		fmt.Printf("close %s\n", me.filename)
	} else if nil != me.w {
		me.w.Close()
		fmt.Printf("close %s\n", me.filename)
	}
}

func hdfsClose() {
	for _, st := range sts {
		if nil != st {
			st.close()
		}
	}

	if nil != client {
		//client.Close()
		fmt.Printf("client close\n")
	}
}

func initCache() {
	if true {
		seed := rand.New(rand.NewSource(time.Now().UnixNano()))
		fmt.Println("init ...")
		seed.Read(cache[:])
		fmt.Println("init ok.")
	}
}

func help() {
	usage := fmt.Sprintf(": %s {read|write} FILENAME COUNT", os.Args[0])

	fmt.Fprintln(os.Stderr, usage)
}

func main() {
	namenode = os.Getenv("NAMENODE")
	fmt.Println("NAMENODE =", namenode)
	if "" == namenode {
		help()

		return
	}

	username = os.Getenv("USERNAME")
	fmt.Println("USERNAME =", username)
	if "" == username {
		help()

		return
	}

	if len(os.Args) < 4 {
		help()

		return
	}

	client, _ = hdfs.NewForUser(namenode, username)

	// db, _ = leveldb.OpenFile("/home/liujf/db", nil)
	// defer db.Close()

	op := os.Args[1]
	filename := os.Args[2]
	count, _ := strconv.Atoi(os.Args[3])
	sts = make([]*hdfsst, count)

	ch := make(chan os.Signal, 0)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go do_exit(ch)

	begin = time.Now()

	switch op {
	case "read":
		//		batch_read(filename, count)
		read(filename, 0)
	case "write":
		//		batch_write(filename, count)
		write(filename, 0)
	default:
		help()
	}

	//	wg.Wait()
	do_statistics()
}

func batch_read(filename string, count int) {
	for i := 0; i < count; i++ {
		//		wg.Add(1)
		go read(filename, i)
	}
}

func batch_write(filename string, count int) {
	initCache()

	for i := 0; i < count; i++ {
		//		wg.Add(1)
		go write(filename, i)
	}
}

func read(filename string, idx int) {
	//	defer wg.Add(-1)

	offset := int64(0)
	ifilename := filename + "-" + strconv.Itoa(idx)
	st := &hdfsst{
		filename: ifilename,
	}
	sts[idx] = st

	info, err := client.Stat(filename)
	if nil != err {
		fmt.Printf("stat %s error:%v\n", ifilename, err)

		return
	}
	size := info.Size()
	fmt.Printf("stat %s size:%d\n", ifilename, size)

	file, err := client.Open(filename)
	if nil != err {
		fmt.Printf("open %s error:%v\n", ifilename, err)

		return
	}
	st.r = file

	fmt.Printf("open %s\n", ifilename)
	fmt.Printf("read %s ...\n", ifilename)

	for {
		for _, f := range frags {
			times := 0

			for i := 0; i < f.count; i++ {
				if offset+int64(f.size) >= size {
					goto done
				}

				file.Seek(offset, 0)
				file.Read(cache[:f.size])
				times++
				if 0 == i%f.times {
					fmt.Printf("read %s offset:%d size:%d times:%d time:%v\n", ifilename, offset, f.size, times, time.Now())
				}
				offset += int64(f.size)
				st.statistics(f.size)
			}
		}
	}

done:
	fmt.Printf("read %s ok.\n", ifilename)
}

func write(filename string, idx int) {
	//	defer wg.Add(-1)

	filename += "-" + strconv.Itoa(idx)
	st := &hdfsst{
		filename: filename,
	}
	sts[idx] = st

	if err := client.CreateEmptyFile(filename); nil == err {
		fmt.Printf("create %s\n", filename)
	}

	file, err := client.Append(filename)
	if nil != err {
		fmt.Printf("open %s error:%v\n", filename, err)
		return
	} else {
		fmt.Printf("open %s\n", filename)
	}
	st.w = file

	fmt.Printf("write %s ...\n", filename)

	for _, f := range frags {
		times := 0

		for i := 0; i < f.count; i++ {
			times++
			file.Write(cache[:f.size])

			if 0 == i%f.times {
				fmt.Printf("append %s size:%d times:%d time:%v\n", filename, f.size, times, time.Now())
			}
			st.statistics(f.size)
		}
	}

	fmt.Printf("write %s ok.\n", filename)
}

func do_statistics() {
	d := time.Now().Sub(begin)

	second := d.Seconds()

	total := int64(0)
	for idx, st := range sts {
		total += st.bytes

		fmt.Printf("st%d times:%d\n", idx, st.times)
	}

	t := float64(total) / float64(M)

	fmt.Printf("total:%fM, second:%fs, throughput:%fM/s\n", t, second, t/second)

	hdfsClose()
	os.Exit(0)
}

func do_exit(ch chan os.Signal) {
	s := <-ch
	fmt.Printf("recv signal %s\n", s)

	do_statistics()
}
