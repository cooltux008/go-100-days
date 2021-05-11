package main

import . "fmt"

func change(val *int) {
	*val = 55
}

func main() {
	a := 58
	Println("value of a before function call is", a)

	b := &a
	change(b)
	Println("value of a after function call is", a)
}
