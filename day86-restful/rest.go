package main

import (
	"context"
	"fmt"
	"go-100-days/day86-restful/student"
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
)

type Student struct {
}

var (
	cli student.StudentService
)

func (s *Student) GetStudent(req *restful.Request, resp *restful.Response) {
	name := req.PathParameter("name")
	fmt.Println(name)

	response, err := cli.GetStudent(context.TODO(), &student.Request{
		Name: name,
	})
	if err != nil {
		fmt.Println(err.Error())
		resp.WriteError(500, err)
	}
	resp.WriteEntity(response)
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.student"),
	)

	service.Init()

	cli = student.NewStudentService("go.micro.srv.student", client.DefaultClient)
	student := new(Student)

	ws := new(restful.WebService)
	ws.Path("/student")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/{name}").To(student.GetStudent))
	wc := restful.NewContainer()
	wc.Add(ws)

	service.Handle("/", wc)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
