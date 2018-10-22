package main

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"grpc_demo/helloworld"
)

// client.go

const (
	address     = "localhost:50002"
	defaultName = "hello world"

)

func RunClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	name := defaultName
	r, err := c.SayHello(context.Background(), & helloworld.HelloRequest{Name:name})
	fmt.Println("from server: ", r.Message)
}

func main() {
	RunClient()
}