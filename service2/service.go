// author: ashing
// time: 2020/4/6 9:55 下午
// mail: axingfly@gmail.com
// Less is more.

package service2

import (
	"net/rpc"

	"github.com/ronething/grpc-sample/constant"
)

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(constant.HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}
