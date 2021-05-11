package main

import (
	"fmt"

	"github.com/go-100-days/day13-struct/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}
