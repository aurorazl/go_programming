package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcDemo/proto/package/hello"
	"grpcDemo/utils"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServer
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	if in.Area == pb.Area_B {
		return &pb.HelloResponse{Message: "world " + in.GetName()}, nil
	}
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GetNameList(ctx context.Context, in *pb.QuerySize) (*pb.NameResponseList, error) {
	return &pb.NameResponseList{
		NameRes: []*pb.HelloResponse{
			&pb.HelloResponse{Message: "czl"},
			&pb.HelloResponse{Message: "hello"},
		},
	},nil
}

func (s *server) GetProdInfo(context.Context, *pb.HelloRequest) (*pb.ProdModel, error) {
	return &pb.ProdModel{ProdId: 1, ProdName: "czl", ProdPrice: 21.1},nil
}

func main () {
	lis, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//s := grpc.NewServer(grpc.Creds(utils.GetTls()))
	s := grpc.NewServer(grpc.Creds(utils.GetServerTls()))

	pb.RegisterHelloServer(s, &server{})
	pb.RegisterOrderServiceServer(s, &OrderService{})
	pb.RegisterUserServiceServer(s, &UserService{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}