package main

import (
	"context"
	"errors"
	"fmt"
	"go-100-days/day86-restful/student"
	"log"

	"github.com/micro/go-micro"
)

type StudentServiceImpl struct {
}

func (ss *StudentServiceImpl) GetStudent(ctx context.Context, request *student.Request, resp *student.Student) error {
	studentMap := map[string]student.Student{
		"dave":   student.Student{Name: "dave", Classes: "软件工程专业", Grade: 80},
		"steven": student.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   student.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   student.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New("请求参数错误，请重新请求。")
	}

	student := studentMap[request.Name]
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*resp = student
		return nil
	}

	return errors.New("未查询到相当学生信息")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.student"),
	)
	service.Init()

	student.RegisterStudentServiceHandler(service.Server(), new(StudentServiceImpl))
	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
