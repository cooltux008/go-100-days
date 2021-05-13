package main

import (
	"fmt"
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileNmae := "/tmp/bar.txt"
	file, err := os.OpenFile(fileNmae, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()

	bs := []byte{97, 98, 99, 100}

	n, err := file.Write(bs)
	fmt.Println(n)
	HandleErr(err)

	file.WriteString("\n")

	//直接写出字符串
	n, err = file.WriteString("Hello World")
	fmt.Println(n)
	HandleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)
}
