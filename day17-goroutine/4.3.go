package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10

var wg sync.WaitGroup
var matex sync.Mutex

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		matex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			matex.Unlock()
			fmt.Println(name, "没票了")
			break
		}
		matex.Unlock()
	}
}

func main() {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")
	wg.Wait()
	//	time.Sleep(5 * time.Second)

}
