package main

import (
	"fmt"
)

func main() {
	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(a)

	fmt.Println(a[1])
	fmt.Println(a[1][len(a[1])-1])
}
