package main

import (
"flag"
"net/http"

"github.com/golang/glog"
"github.com/grpc-ecosystem/grpc-gateway/runtime"
"golang.org/x/net/context"
"google.golang.org/grpc"

	"grpc_demo/helloworld"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:50002", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := helloworld.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}