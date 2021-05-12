package main

import (
	"fmt"
	"strconv"
)

type my_fun func(int, int) string

func fun1() my_fun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

func main() {
	res1 := fun1()
	fmt.Println(res1(10, 20))
}
