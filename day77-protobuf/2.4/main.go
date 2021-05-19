package main

import (
	"fmt"
	"os"

	"go-100-days/day77-protobuf/2.4/person"

	"github.com/golang/protobuf/proto"
)

func main() {
	msgTest := &person.Person{
		Name: proto.String("foo"),
		Age:  proto.Int(18),
		From: proto.String("China"),
	}

	msgDataEncoding, err := proto.Marshal(msgTest)
	if err != nil {
		panic(err.Error())
		return
	}

	msgEntity := person.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("姓名: %s\n", msgEntity.GetName())
	fmt.Printf("年龄: %d\n", msgEntity.GetAge())
	fmt.Printf("国籍: %s\n", msgEntity.GetFrom())
}
