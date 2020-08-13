
package hello

import (
	"context"
    
    "jupiter-test/pb/hello"
    
)

type HelloServiceServer struct{}

func (server *HelloServiceServer) SayHello(context context.Context, request *hello.HelloRequest) (response *hello.HelloResponse, err error) {
	panic("implement me")
    return 
}

