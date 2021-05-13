package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("/tmp/5.1.txt", os.O_RDWR, 0)
	defer file.Close()

	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))
	file.Seek(2, 0)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(3, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd)
	file.WriteString("ABC")
}
