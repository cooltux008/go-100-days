package main

import (
	"fmt"
)

func main() {
	var ch1 chan bool
	fmt.Println(ch1)
	fmt.Printf("%T\n", ch1)

	ch1 = make(chan bool)
	fmt.Println(ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("child goroutine, i:", i)
		}
		ch1 <- true
		fmt.Println("end")
	}()
	data := <-ch1
	fmt.Println("data-->", data)
	fmt.Println("main over...")
}
