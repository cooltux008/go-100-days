package main

import (
	"go-100-days/day82/clientstream/message"

	"google.golang.org/grpc"

	"fmt"
	"io"
	"net"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println("客户端流RPC模式")

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			result := message.OrderInfo{OrderStatus: "读取数据结束"}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(orderRequest.GetOrderId(), orderRequest.GetOrderName(), orderRequest.GetOrderStatus())
	}

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
