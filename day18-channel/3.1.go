package main

import (
	"fmt"
	"time"
)

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func main() {
	ch1 := make(chan int)
	go sendData(ch1)

	for {
		time.Sleep(1 * time.Second)
		v, ok := <-ch1
		if !ok {
			fmt.Println("已经读取了所有的数据, ", ok)
			break
		}
		fmt.Println("取出数据: ", v, ok)

	}
	fmt.Println("main...over...")
}
