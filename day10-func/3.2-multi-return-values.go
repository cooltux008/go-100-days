package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func sumAndMultipied(A, B int) (sum, multiplied int) {
	sum = A + B
	multiplied = A * B
	return sum, multiplied
}
func main() {
	a, b := swap("Mahesh", "kumar")
	fmt.Println(a, b)

	fmt.Println(sumAndMultipied(3, 5))
}
