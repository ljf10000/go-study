package main

import (
	. "asdf"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("hexdump")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		Log.Error("get stdin error:%s", err)
	}

	go func() {
		s := "exec command"
		stdin.Write([]byte(s))
		stdin.Close()
	}()

	err = cmd.Start()
	if nil != err {
		Log.Error("cmd start error:%s", err)
	}

	time.Sleep(3 * time.Second)
}
