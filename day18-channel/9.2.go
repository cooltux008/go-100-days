package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------------------")

	timer2 := time.NewTimer(5 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 结束")
	}()

	time.Sleep(3 * time.Second)
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2  停止")
	}
}
