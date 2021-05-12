package main

import (
	"fmt"
)

type Human interface {
	Len()
}

type Student interface {
	Human
}

type Test struct {
}

func (h *Test) Len() {
	fmt.Println("Success")
}

func main() {
	var s Student
	s = new(Test)
	s.Len()
}
