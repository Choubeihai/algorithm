package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**

标准库的RPC默认采用Go语言特有的gob编码，因此从其它语言调用Go语言实现的RPC服务将比较困难。
在互联网的微服务时代，每个RPC以及服务的使用者都可能采用不同的编程语言，因此跨语言是互联网时代RPC的一个首要条件。
得益于RPC的框架设计，Go语言的RPC其实也是很容易实现跨语言支持的。

Go语言的RPC框架有两个比较有特色的设计：
1. rpc数据打包时通过自定义插件实现自定义的编码和解码，默认是golang的gob编解码
2. rpc建立在io.ReadWriteCloser接口之上，因此可以将rpc架在不同的通讯协议之上

这里尝试通过官方自带的net/rpc/jsonrpc扩展实现一个跨语言的RPC
*/

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
		// rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
