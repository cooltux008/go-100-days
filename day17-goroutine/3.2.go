package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT-->", runtime.GOROOT())
	fmt.Println("/os/platform-->", runtime.GOOS)

	fmt.Println("逻辑CPU的核数", runtime.NumCPU())
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine...")
		}
	}()

	for i := 0; i < 4; i++ {
		runtime.Gosched()
		fmt.Println("main...")
	}
}
