package main

import (
	"fmt"
)

type data [2]int

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)
		fallthrough
	case 7:
		x += 1
		fmt.Println("fallthrough ", x)
	}

}
