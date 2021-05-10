package main

import (
	"fmt"
)

func main() {
	var x interface{}
	x = 10.0
	switch i := x.(type) {
	case nil:
		fmt.Printf("x => %T", i)
	case int:
		fmt.Printf("x => int")
	case float64:
		fmt.Printf("x => float64")
	case func(int) float64:
		fmt.Printf("x => func(int)")
	case bool, string:
		fmt.Printf("x => bool or string")
	default:
		fmt.Printf("unknow")
	}
}
