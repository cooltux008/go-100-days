package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

type AddParam struct {
	Arg1 float32
	Arg2 float32
}

func (mu *MathUtil) Add(param AddParam, resp *float32) error {
	*resp = param.Arg1 + param.Arg2
	return nil
}

func main() {
	mathUtil := new(MathUtil)

	err := rpc.RegisterName("MathUtil", mathUtil)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	http.Serve(listen, nil)
}
