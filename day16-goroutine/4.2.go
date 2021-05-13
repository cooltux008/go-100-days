package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 10

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			fmt.Println(name, "没票了")
			break
		}
	}
}

func main() {
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(5 * time.Second)

}
