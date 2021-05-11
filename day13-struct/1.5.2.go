package main

import "fmt"

type Address struct {
	city, state string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main() {

	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}

	fmt.Println("name: ", p.name)
	fmt.Println("age: ", p.age)
	fmt.Println("city: ", p.address.city)
	fmt.Println("state: ", p.address.state)

}
