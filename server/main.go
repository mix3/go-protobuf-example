package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/mix3/protobuf-example/helloworld"
	"golang.org/x/net/context"
)

type server struct{}

func (s server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
