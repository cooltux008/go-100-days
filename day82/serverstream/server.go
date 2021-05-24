package main

import (
	"go-100-days/day82/serverstream/message"
	"time"

	"google.golang.org/grpc"

	"fmt"
	"net"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println("服务端流RPC模式")

	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "己付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "衣服", OrderStatus: "己付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "衣服", OrderStatus: "己付款"},
	}

	for id, info := range orderMap {
		if time.Now().Unix() >= request.TimeStamp {
			fmt.Println("订单序列号ID: ", id)
			fmt.Println("订单详情: ", info)
			stream.Send(&info)
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	server.Serve(listen)
}
