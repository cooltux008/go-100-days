package main

import (
	"fmt"
	"go-100-days/day80/3.1/message"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907310001", TimeStamp: timeStamp}

	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*response)
}
