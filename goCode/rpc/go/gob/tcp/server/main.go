package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	rpc.RegisterName("helloServiceName", new(HelloService))
	listener, _ := net.Listen("tcp", ":9092")
	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}
