package main

import (
	"go-100-days/day80/3.1/message"
	"time"

	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "己付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "衣服", OrderStatus: "己付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "衣服", OrderStatus: "己付款"},
	}

	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		fmt.Println(result)
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		} else {
			fmt.Println(request.OrderId)
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderService := new(OrderService)

	rpc.Register(orderService)
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	http.Serve(listen, nil)
}
