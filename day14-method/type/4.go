package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) Show() {
	fmt.Println("Person-->", p.name)
}

type People = Person

type Student struct {
	Person
	People
}

func (p People) Show2() {
	fmt.Println("People------>", p.name)
}

func main() {
	var s Student

	s.People.name = "lixiaohua"
	s.Person.name = "wangergou"
	s.Person.Show()
	s.People.Show2()
	fmt.Printf("%T,%T\n", s.Person, s.People)
}
