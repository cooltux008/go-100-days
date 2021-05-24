package main

import (
	"go-100-days/day82/clientserverstream/message"

	"google.golang.org/grpc"

	"fmt"
	"io"
	"net"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
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
