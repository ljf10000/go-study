package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	. "asdf"
)

var conf string
var app App
var wg sync.WaitGroup

type App struct {
	PidFile  FileName
	StatPath string
	Monitors []*Monitor
}

type Monitor struct {
	Name     string
	PidFile  FileName
	Cmd      string
	Pid      int
	Times    int
	Error    int
	Create   string
	StatFile FileName

	path    string
	argv    []string
	process *os.Process
}

func (me *Monitor) isDeamon() bool {
	return Empty != me.PidFile
}

func (me *Monitor) init() {
	argv := strings.Split(me.Cmd, " ")
	me.path = argv[0]
	argv[0] = path.Base(argv[0])
	me.argv = argv

	if Empty == me.Name {
		me.Name = me.path
	}

	me.StatFile = FileName(app.StatPath + "/" + me.Name + "-monitor.stat")
}

func (me *Monitor) update() {
	me.Times++
	me.Create = time.Now().String()

	j := "{}"
	b, err := json.Marshal(me)
	if nil == err {
		j = string(b)
	}

	me.StatFile.Saves([]string{j}, true)
}

func (me *Monitor) readPid() int {
	return me.PidFile.ReadPid()
}

func (me *Monitor) runNormal() {
	defer me.update()

	process, err := os.StartProcess(me.path, me.argv, StdAttr)
	if nil != err {
		Log.Info("start normal %v error:%v", me.argv, err)

		me.Error++

		return
	}

	me.process = process
	me.Pid = process.Pid

	Log.Info("start normal %v pid %v", me.argv, me.Pid)

	ps, err := me.process.Wait()
	if nil != err {
		Log.Info("wait normal %s:%d error:%v", me.path, me.Pid, err)
	} else if ps.Exited() {
		Log.Info("normal %s:%d exit", me.path, me.Pid)
	}

	me.process.Release()
}

func (me *Monitor) runDeamon() {
	_, err := os.StartProcess(me.path, me.argv, StdAttr)
	if nil != err {
		Log.Info("start deamon %v error:%v", me.argv, err)

		me.Error++
		me.update()

		return
	}

	Log.Info("start deamon %v", me.argv)
	// first get pid
	for {
		time.Sleep(100 * time.Millisecond)

		pid := me.readPid()
		if pid > 0 {
			me.Pid = pid
			me.update()

			break
		}
	}

	defer func() {
		me.Pid = -1

		Log.Info("deamon %s exit", me.path)
	}()

	for {
		// get pid again
		pid := me.readPid()
		if pid != me.Pid {
			Log.Info("pid:%v != me.pid:%v", pid, me.Pid)

			return
		}

		p, err := os.FindProcess(pid)
		if nil != err {
			Log.Info("find process %d error:%v", pid, err)

			return
		}

		ps, err := p.Wait()
		if nil == err && ps.Exited() {
			p.Release()

			return
		}

		time.Sleep(3 * time.Second)
		Log.Info("deamon %s EXIST", me.path)
	}
}

func (me *Monitor) run() {
	if me.isDeamon() {
		me.runDeamon()
	} else {
		me.runNormal()
	}
}

func init() {
	flag.StringVar(&conf, "conf", Empty, "config file")
}

func load() {
	flag.Parse()

	if Empty == conf {
		os.Exit(-1)
	}

	f, err := os.Open(conf)
	defer f.Close()
	if nil != err {
		Log.Error("open %s error:%v", conf, err)

		os.Exit(-2)
	}

	buf, err := ioutil.ReadAll(f)
	if nil != err {
		Log.Error("read %s error:%v", conf, err)

		os.Exit(-3)
	}

	if err := json.Unmarshal(buf, &app); nil != err {
		Log.Info("invalid json config:%s error:%v", conf, err)

		os.Exit(-4)
	}

	if 0 == len(app.Monitors) {
		Log.Info("monitor nothing, exit.")

		os.Exit(-5)
	} else if Empty == app.PidFile {
		Log.Info("no pid file.")

		os.Exit(-6)
	}

	/*
		pid := readPidFile(app.PidFile)
		if process, err := os.FindProcess(pid); nil == err {
			process.Pid
			Log.Info("EXIT EXIT EXIT EXIT EXIT")

			os.Exit(-7)
		}
	*/

	if Empty == app.StatPath {
		app.StatPath = "/tmp"
	}

	app.PidFile.WritePid()
}

func main() {
	load()

	for _, m := range app.Monitors {
		m.init()
		wg.Add(1)

		go func(m *Monitor) {
			Log.Info("start %s", m.path)

			for {
				m.run()
			}

			wg.Add(-1)
		}(m)
	}

	wg.Wait()
}
