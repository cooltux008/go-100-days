package main

import (
	"fmt"

	"github.com/go-100-days/day14-method/interface/1.3/test"
)

type T struct {
	test.Controller
}

func (t *T) Get() {
	fmt.Print("T")
}

func (t *T) Post() {
	fmt.Print("T")
}

func main() {
	var something test.Something
	something = new(T)

	var t T
	t.M = 1
	something.Get()
}
