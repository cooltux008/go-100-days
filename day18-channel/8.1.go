package main

import (
	"fmt"
)

func func1(ch chan<- int) {
	ch <- 100
	fmt.Println("func1函数结束")
}

func func2(ch <-chan int) {
	data := <-ch
	fmt.Println("func2函数，从ch中读取的数据是: ", data)
}

func main() {
	ch1 := make(chan int)
	go func1(ch1)
	data := <-ch1
	fmt.Println("func1中写出的数据是: ", data)

	go func2(ch1)
	ch1 <- 200
	fmt.Println("main...over...")
}
