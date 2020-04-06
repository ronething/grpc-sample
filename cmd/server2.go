// author: ashing
// time: 2020/4/6 10:04 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ronething/grpc-sample/service2"
)

func main() {
	service2.RegisterHelloService(new(service2.HelloService))

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	log.Printf("server on port: 1234\n")
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
