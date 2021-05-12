package main

import (
	"fmt"
)

func main() {
	var i1 interface{} = new(Student)
	s := i1.(Student)

	fmt.Println(s)

	var i2 interface{} = new(Student)
	s, ok := i2.(Student)
	if ok {
		fmt.Println(s)
	}
}

type Student struct {
}
