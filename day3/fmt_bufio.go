package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input a string:")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('.')
	fmt.Println("the data: ", s1)
}
