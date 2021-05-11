package main

import (
	"fmt"
)

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	for _, v := range []rune(name) {
		fmt.Printf("%c", v)
	}
	fmt.Printf("\nReversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

}
