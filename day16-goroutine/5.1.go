package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func fun1() {
	for i := 1; i <= 10; i++ {
		fmt.Println("fun1...i:", i)
	}
	wg.Done()
}
func fun2() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("\tfun2...i:", i)
	}
}

func main() {
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态，等待wg中的child goroutine结束")
	wg.Wait()
	fmt.Println("main, 解除阻塞")

}
