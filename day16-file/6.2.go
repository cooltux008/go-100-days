package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//读取文件中的所有数据
	fileName1 := "/tmp/aa.txt"
	data, err := ioutil.ReadFile(fileName1)
	fmt.Println(err)
	fmt.Println(string(data))

	//写出数据
	fileName2 := "/tmp/bb.txt"
	s1 := "helloworld我爱中国"
	err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
	fmt.Println(err)

	//ReadAll
	s2 := "qwertyuiopsdfghjklzxcvbnm"
	r1 := strings.NewReader(s2)
	data, _ = ioutil.ReadAll(r1)
	fmt.Println(data)

	//ReadDir
	dirName := "/tmp/a"
	fileInfos, _ := ioutil.ReadDir(dirName)
	fmt.Println(len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		fmt.Printf("%%\n", fileInfos[i])
		fmt.Println(i, fileInfos[i].Name(), fileInfos[i].IsDir())
	}

	//TempDir
	dir, err := ioutil.TempDir("/tmp/temp", "Test")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(dir)
	fmt.Printf("%s\n", dir)
	//TempFile
	f, err := ioutil.TempFile("/tmp/tempfile", "Test")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(f.Name())
	fmt.Printf("%s\n", f.Name())
}
