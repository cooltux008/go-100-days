package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	args := os.Args
	dirname := args[1]
	listFiles(dirname, 0)
}

func listFiles(dirname string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := path.Join(dirname, fi.Name())
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			listFiles(filename, level+1)
		}
	}
}
