package main

import (
	"fmt"
)

func test1(ch chan int) {
	fmt.Printf("%T, %p\n", ch, ch)
}

func main() {
	ch1 := make(chan int)
	fmt.Printf("%T,%p\n", ch1, ch1)
	test1(ch1)
}
