package main

import . "fmt"

func main() {
	var a int = 20
	Printf("a变量的地址：%x\n", &a)

	var ip *int
	ip = &a
	Printf("ip变量的地址：%x\n", ip)

	Printf("*ip 变量的值：%d\n", *ip)
}
