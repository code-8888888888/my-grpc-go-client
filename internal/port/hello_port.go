package port

import (
	"context"

	"github.com/code-8888888888/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type HelloClientPort interface {
	SayHello(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (*hello.HelloResponse, error)
	SayManyHellos(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (hello.HelloService_SayManyHellosClient, error)
	SayHelloToEveryone(ctx context.Context, opts ...grpc.CallOption) (hello.HelloService_SayHelloToEveryoneClient, error)
	SayHelloContinuos(ctx context.Context, opts ...grpc.CallOption) (hello.HelloService_SayHelloContinuosClient, error)
}