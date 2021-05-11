package main

import . "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	Printf("变量a = %d\n", a)
	Printf("指针变量*ptr = %d\n", *ptr)
	Printf("指向指针的指针变量**pptr = %d\n", **pptr)
}
