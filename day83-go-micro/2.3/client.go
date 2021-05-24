package main

import (
	"context"
	"fmt"
	"go-100-days/day83-go-micro/2.3/message"
	"time"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("student.client"))
	service.Init()

	studentService := message.NewStudentService("student_service", service.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "dave"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)

	time.Sleep(50 * time.Second)
}
