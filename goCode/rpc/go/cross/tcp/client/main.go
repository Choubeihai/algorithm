package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	var reply string
	conn, _ := net.Dial("tcp", "82.157.238.131:9092")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	if err := client.Call("helloServiceName.Hello", "caixiaowei", &reply); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}
