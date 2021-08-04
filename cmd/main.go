package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func found(ss []string) bool {
	for i := 0; i < len(ss); i++ {
		if ss[i] == ".0" {
			return true
		}
	}

	return false
}

func main() {
	args := "snmpvsm"
	for i := 1; i < len(os.Args); i++ {
		args += " '" + os.Args[i] + "'"
	}
	args += " 2>&1"

	cmd := exec.Command("/bin/bash", "-c", args)

	stdout, err := cmd.StdoutPipe()
	if nil != err {
		os.Exit(1)
	}

	cmd.Start()

	stdout_reader := bufio.NewReader(stdout)

	for {
		line, err := stdout_reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		s := strings.TrimSuffix(line, "\n")
		s = strings.Trim(s, " ")
		if s == "." || s == "./" || s == ".." || s == "../" {
			fmt.Printf(line)
		} else if ss := strings.Split(s, "/"); false == found(ss) {
			fmt.Printf(line)
		}
	}

	os.Exit(0)
}
