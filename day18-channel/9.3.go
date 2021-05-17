package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch1)

	fmt.Println(time.Now())
	time2 := <-ch1
	fmt.Println(time2)
}
