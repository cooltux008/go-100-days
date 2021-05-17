package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("name: %s, age: %n, gender: %s\n", p.Name, p.Age, p.Sex)
}

func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func main() {
	p2 := Person{"Ruby", 30, "male"}

	getValue := reflect.ValueOf(p2)
	methodValue1 := getValue.MethodByName("PrintInfo")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue1.Kind(), methodValue1.Type())
	methodValue1.Call(nil)

	args1 := make([]reflect.Value, 0)
	methodValue1.Call(args1)

	methodValue2 := getValue.MethodByName("Say")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue2.Kind(), methodValue2.Type())
	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := getValue.MethodByName("Test")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue3.Kind(), methodValue3.Type())
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("Hello")}
	methodValue3.Call(args3)
}
