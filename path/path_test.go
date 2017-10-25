package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test1(t *testing.T) {
	filepath.Walk("D:\\code\\lang\\go\\gopath\\src", func(path string, f os.FileInfo, err error) error {
		if nil == f {
			return err
		} else if f.IsDir() {
			return nil
		}

		fmt.Println("path=", path)

		return nil
	})
}
