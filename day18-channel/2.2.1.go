package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	done := make(chan bool)
	go func() {
		fmt.Println("child goroutine")
		time.Sleep(3 * time.Second)
		data := <-ch1
		fmt.Println("data: ", data)
		done <- true
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 100

	<-done
	fmt.Println("main over")
}
