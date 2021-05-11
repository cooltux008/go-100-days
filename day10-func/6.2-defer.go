package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	defer fmt.Println(b)
	fmt.Println(a)
}
