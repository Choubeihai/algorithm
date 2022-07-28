package main

import (
	"fmt"
	"reflect"
)

func main() {
	person := NewUseReflect().(*People)
	fmt.Println(person)
}

type People struct {
	Age  int
	Name string
}

func New() *People {
	return &People{
		Age:  18,
		Name: "shiina",
	}
}

func NewUseReflect() interface{} {
	var p People
	t := reflect.TypeOf(p)
	v := reflect.New(t)
	v.Elem().Field(0).Set(reflect.ValueOf(18))
	v.Elem().Field(1).Set(reflect.ValueOf("shiina"))
	return v.Interface()
}
