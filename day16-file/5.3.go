package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "/tmp/english.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	//创建Reader对象
	b1 := bufio.NewReader(file)
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	fmt.Println(string(p[:n1]))

	//ReadLine()
	data, flag, err := b1.ReadLine()
	fmt.Println(flag)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	//ReadString()
	s1, err := b1.ReadString('\n')
	fmt.Println(err)
	fmt.Println(s1)

	for {
		s1, err := b1.ReadSlice('\n')
		if err == io.EOF {
			fmt.Println("读取完毕。。")
			break
		}
		fmt.Println(s1)
	}

	//ReadBytes()
	data, err = b1.ReadBytes('\n')
	fmt.Println(err)
	fmt.Println(string(data))

	//Scanner
	s2 := ""
	fmt.Scanln(&s2)
	fmt.Println(s2)

	b2 := bufio.NewReader(os.Stdin)
	s2, _ = b2.ReadString('\n')
	fmt.Println(s2)
}
