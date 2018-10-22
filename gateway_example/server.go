package main

// server.go
import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc_demo/helloworld"
	"net"
)

const (
	port = ":50002"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("from client: ", in.Name)
	return &helloworld.HelloReply{Message: in.Name}, nil
}

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

func main() {
	Run()
}
