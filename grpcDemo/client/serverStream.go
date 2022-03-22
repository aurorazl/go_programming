package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/proto/package/hello"
	"grpcDemo/utils"
	"io"
	"log"
	"time"
)


var (
	addr = flag.String("addr", "localhost:8888", "the address to connect to")
	name = flag.String("name", "defaultName", "Name to greet")
)

func main() {

	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(getTls()))
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(utils.GetClientTls()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//client := pb.NewHelloClient(conn)
	client := pb.NewUserServiceClient(conn)

	stream, err := client.ReturnUsersByStream(ctx, &pb.UserRequest{Users: []*pb.UserInfo{
		&pb.UserInfo{UserId: 1, UserName: "czl"},
		&pb.UserInfo{UserId: 2, UserName: "czl1"},
		&pb.UserInfo{UserId: 3, UserName: "czl2"},
		&pb.UserInfo{UserId: 4, UserName: "3"},
		&pb.UserInfo{UserId: 5, UserName: "czl4"},
		&pb.UserInfo{UserId: 6, UserName: "czl5"},
	}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		userResponse, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userResponse)
	}

}
