package main

import (
	"fmt"
)

type name struct {
	firstName string
	lastName  string
}

func main() {
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}

	if name1 == name2 {
		fmt.Println("name1 == name2")
	} else {
		fmt.Println("name1 != name2")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 == name4")
	} else {
		fmt.Println("name3 != name4")
	}

}
