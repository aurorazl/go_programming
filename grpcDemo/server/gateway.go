package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpcDemo/utils"
	pb "grpcDemo/proto/package/hello"
	"log"
	"net/http"
)

func gatewayService() {
	mux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(utils.GetClientTls())}
	err := pb.RegisterHelloHandlerFromEndpoint(context.Background(), mux, "localhost:8888", opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}
