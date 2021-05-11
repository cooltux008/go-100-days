package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64
	fmt.Println("input a int & a float64:")
	fmt.Scanln(&x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d, y:%f\n", x, y)
}
