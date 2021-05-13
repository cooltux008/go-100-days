package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "/tmp/foo.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()

	bs := make([]byte, 4, 4)
	//1
	n, err := file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	//2
	n, err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	//3
	n, err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	//4
	n, err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	n = -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
