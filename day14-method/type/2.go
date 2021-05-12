package main

import (
	"fmt"
)

type myint int    //定义新类型
type myint2 = int //不是定义新类型，仅仅起别名

func main() {
	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1, i2)

	var i3 myint2
	i3 = i2
	fmt.Println(i1, i2, i3)

	var he = "hello"
	fmt.Println(he)
}
