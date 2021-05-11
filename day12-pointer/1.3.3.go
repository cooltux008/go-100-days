package main

import . "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

func main() {
	var a = first{1, false, 2}
	var b *first = &a
	Println(a.b, a.a, a.name, &a, b.a, &b, (*b).a)
}
