package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func writeData(i int) {
	defer wg.Done()
	rwMutex.Lock()
	fmt.Println(i, "正在写")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	//fmt.Println(i, "写结束")
}

func readData(i int) {
	defer wg.Done()

	rwMutex.RLock()
	fmt.Println(i, "正在读")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock()
	//	fmt.Println(i, "读结束")
}

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	wg.Add(11)
	go writeData(1)
	for i := 2; i <= 10; i++ {
		go readData(i)
	}
	go writeData(11)
	wg.Wait()
	fmt.Println("main...over...")
}
