package main

import (
	"context"
	"fmt"
	"go-100-days/day81/2.1/message"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	orderRequest := &message.OrderRequest{OrderId: "201907310001", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
