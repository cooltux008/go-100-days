package main

import (
	"errors"
	"fmt"
)

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
		}
	}()
	return funcB()
}

func funcB() error {
	panic("foo")
	return errors.New("success")
}

func main() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}
