package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arg := os.Args[1]
	files, error := filepath.Glob(arg)
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}
