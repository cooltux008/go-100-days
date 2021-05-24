package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go-100-days/day83-go-micro/2.3/message"

	"github.com/micro/go-micro"
)

type StudentManager struct {
}

func (sm *StudentManager) GetStudent(ctx context.Context, request *message.StudentRequest, response *message.Student) error {
	//tom
	studentMap := map[string]message.Student{
		"dave":   message.Student{Name: "dave", Classes: "软件工程专业", Grade: 80},
		"steven": message.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   message.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   message.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New("请求参数错误，请重新请求")
	}

	student := studentMap[request.Name]
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*response = student
		return nil
	}

	return errors.New("未查询到相关学生信息")
}

func main() {
	service := micro.NewService(micro.Name("student_service"), micro.Version("v1.0.0"))
	service.Init()

	message.RegisterStudentServiceHandler(service.Server(), new(StudentManager))

	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
