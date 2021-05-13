package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileName1 := "/Users/admin/.bash_history"
	fileName2 := "bb.txt"

	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))

	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println("获取父目录", path.Join(fileName1, ".."))

	//create dir
	err := os.Mkdir("/tmp/hello", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("create dir success")
	err = os.MkdirAll("/tmp/foo/a/b/c/d", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("create multi dirs success")

	//create file
	file1, err := os.Create("/tmp/hello/foo.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file1)

	file2, err := os.Create(fileName2)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file2)

	//open file
	file3, err := os.Open(fileName1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file3)

	file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)
	file4.Close()

	//rm file or dir
	err = os.Remove("/tmp/hello/foo.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("rm file success")

	err = os.RemoveAll("/tmp/hello")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("rm dir success")

}
