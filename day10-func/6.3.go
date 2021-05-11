package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

func (p person) age() {
	fmt.Printf("Age: %d \n", 10)
}

func (p person) last() {
	fmt.Println("last")
}

func (p person) begin() {
	fmt.Println("begin")
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "smith",
	}
	defer p.begin()
	defer p.last()

	defer p.age()

	fmt.Printf("Welcome ")
	defer p.fullName()
}
