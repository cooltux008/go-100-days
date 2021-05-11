package main

import (
	"fmt"
)

type image struct {
	data map[int]int
}

func main() {
	image1 := image{data: map[int]int{0: 155}}
	image2 := image{data: map[int]int{0: 155}}

	if image1 == image2 {
		fmt.Println("image1 == image2")
	} else {
		fmt.Println("image1 != image2")
	}

}
