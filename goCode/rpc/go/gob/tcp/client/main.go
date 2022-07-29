package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var reply string
	client, _ := rpc.Dial("tcp", ":9092")
	if err := client.Call("HelloService1.Hello", "caixiaowei", &reply); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}
