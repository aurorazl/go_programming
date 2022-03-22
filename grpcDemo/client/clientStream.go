package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/proto/package/hello"
	"grpcDemo/utils"
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

	stream, err := client.SendUsersByStream(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	var users []*pb.UserInfo
	for i:=0;i<3;i++{
		for j:=0;j<2;j++ {
			users = append(users, &pb.UserInfo{UserId: int32(i+j),UserName: "czl"})
		}
		err = stream.Send(&pb.UserRequest{Users: users})
		users = users[0:0]
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second*1)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Users)
}
