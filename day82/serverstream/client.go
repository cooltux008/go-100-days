package main

import (
	"context"
	"fmt"
	"go-100-days/day82/serverstream/message"
	"time"

	"io"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	orderRequest := &message.OrderRequest{TimeStamp: time.Now().Unix()}
	orderInfosClient, err := orderServiceClient.GetOrderInfos(context.TODO(), orderRequest)

	for {
		orderInfo, err := orderInfosClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("读取到的信息: ", orderInfo)
	}
}
