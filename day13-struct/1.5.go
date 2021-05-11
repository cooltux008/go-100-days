package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println("name: ", mark.name)
	fmt.Println("age: ", mark.age)
	fmt.Println("weight: ", mark.weight)
	fmt.Println("speciality: ", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("Mark changed his speciallity")
	fmt.Println("new speciality: ", mark.speciality)

	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("new age: ", mark.age)

	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("new weight: ", mark.weight)
}
