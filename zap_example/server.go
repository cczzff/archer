package main

// server.go
import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc_demo/helloworld"
	"net"
)

const (
	port = ":50003"
)

type server struct{}

var (
	zapLogger  *zap.Logger
)

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("from client: ", in.Name)
	return &helloworld.HelloReply{Message: in.Name}, nil
}

func Run() {
	zapLogger, _ = zap.NewDevelopment()

	zapLogger.Named("grpc zap")

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return
	}

	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_zap.UnaryServerInterceptor(zapLogger),
	),
	)
	helloworld.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

func main() {
	Run()
}
