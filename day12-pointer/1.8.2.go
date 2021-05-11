package main

import . "fmt"

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func main() {
	var a int = 100
	var b int = 200

	Printf("交换前a的值: %d\n", a)
	Printf("交换前b的值: %d\n", b)
	swap(&a, &b)
	Printf("交换前a的值: %d\n", a)
	Printf("交换前b的值: %d\n", b)

}
