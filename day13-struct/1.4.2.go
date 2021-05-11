package main

import "fmt"

type Books struct {
}

func (s Books) String() string {
	return "data"
}

func main() {
	fmt.Printf("%v\n", Books{})
	mybook := new(Books)
	fmt.Printf(mybook.String())
}
