package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/mix3/protobuf-example/helloworld"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
