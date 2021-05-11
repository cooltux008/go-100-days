package main

import (
	"fmt"
)

func main() {
	var a [4]float32
	fmt.Println(a)

	var b = [5]string{"ruby", "hello", "rose"}
	fmt.Println(b)

	var c = [5]int{'A', 'B', 'C', 'D', 'E'}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	e := [5]int{4: 100}
	fmt.Println(e)

	f := [...]int{0: 1, 4: 1, 9: 1}
	fmt.Println(f)

	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}

	// len
	l := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(l))

	fmt.Println("-------------------")
	for i := 0; i < len(l); i++ {
		fmt.Printf("%d the element of a is %.2f\n", i, l[i])
	}
	fmt.Println("###################")
	sum := float64(0)
	for i, v := range l {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
