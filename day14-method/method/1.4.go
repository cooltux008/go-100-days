package main

import (
	"fmt"
)

type Rectanle struct {
	width, height int
}

func (r *Rectanle) setVal() {
	r.height = 20
}

func main() {
	p := Rectanle{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)

}
