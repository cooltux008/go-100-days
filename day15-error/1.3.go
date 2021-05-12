package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Path)
		fmt.Println(err.Err)
		fmt.Println(err.Op)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
