package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	hello "rpc/go"
)

func main() {
	p := &hello.Person{
		Name:  "晓玮",
		Age:   25,
		Email: "12345",
	}
	bytes, _ := proto.Marshal(p)
	fmt.Println(bytes)

	p1 := &hello.Person{}
	err := proto.Unmarshal(bytes, p1)
	if err != nil {
		return
	}
	fmt.Println(p1)
	fmt.Println(p1.GetAge())
	fmt.Println(p1.GetName())
	fmt.Print(p1.GetEmail())
}
