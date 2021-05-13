package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	srcFile := "/tmp/hello.jpg"
	destFile := "dest.jpg"
	tempFile := destFile + "temp.txt"

	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file1.Close()
	defer file2.Close()

	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	fmt.Println(n1)

	countStr := string(bs[:n1])
	fmt.Println(countStr)

	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	file1.Seek(count, 0)
	file2.Seek(count, 0)
	data := make([]byte, 1024, 1024)
	n2 := -1
	n3 := -1
	total := int(count)

	for {
		n2, err = file1.Read(data)
		if err == io.EOF {
			fmt.Println("文件复制完毕。。")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, _ = file2.Write(data[:n2])
		total += n3

		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		if total > 8000 {
			panic("断电模拟")
		}
	}

}
