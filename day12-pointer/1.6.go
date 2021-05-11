package main

import . "fmt"

func main() {
	b := 255
	a := &b

	Println("address of b is", a)
	Println("value of b is", *a)

	*a++
	Println("new value of b is", b)
}
