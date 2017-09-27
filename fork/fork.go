package main

import (
	"os"
	"os/exec"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := len(os.Args)
	if count < 2 {
		os.Exit(-1)
	}

	cmds := make([]*exec.Cmd, count-1)

	// skip os.Args[0]
	for idx, cmd := range os.Args[1:] {
		argv := strings.SplitN(cmd, " ", 2)

		cmds[idx] = exec.Command(argv[0], argv[1:]...)
	}

	for _, cmd := range cmds {
		wg.Add(1)

		go func(cmd *exec.Cmd) {
			cmd.Run()

			wg.Add(-1)
		}(cmd)
	}

	wg.Wait()
}
