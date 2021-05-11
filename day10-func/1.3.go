package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	fmt.Printf("max: %d\n", ret)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
