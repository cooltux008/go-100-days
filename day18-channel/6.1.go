package main

import (
	"fmt"
	"strconv"
	"time"
)

func sendData3(ch3 chan string) {
	for i := 0; i < 10; i++ {
		ch3 <- "数据" + strconv.Itoa(i)
		fmt.Println("子goroutine, 写出第", i, "个数据")
	}
	close(ch3)
}

func main() {
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))

	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 100
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 200
	fmt.Println("--------------")

	ch3 := make(chan string, 4)
	go sendData3(ch3)
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch3
		if !ok {
			fmt.Println("读完了", ok)
			break
		}
		fmt.Println("\t读取的数据是: ", v)

	}
	fmt.Println("main...over...")
}
