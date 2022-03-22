package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpcDemo/proto/package/hello"
	"grpcDemo/utils"
	"log"
	"time"
)


var (
	addr = flag.String("addr", "localhost:8888", "the address to connect to")
	name = flag.String("name", "defaultName", "Name to greet")
)

func getTls() credentials.TransportCredentials {
	creds, err := credentials.NewClientTLSFromFile("./identity/server.crt", "czl.com")
	if err != nil {
		log.Fatal(err)
	}
	return creds
}

func main() {

	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(getTls()))
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(utils.GetClientTls()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//client := pb.NewHelloClient(conn)
	client := pb.NewOrderServiceClient(conn)

	//r, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name, Area: pb.Area_B})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	//r, err := client.GetNameList(ctx, &pb.QuerySize{Size: 2})
	//log.Printf("Greeting: %s", r.GetNameRes())

	//r, err := client.GetProdInfo(ctx, &pb.HelloRequest{Name: "czl"})

	r, err := client.NewOrder(ctx, &pb.OrderModel{UserId: 1,
		OrderId: 1, OrderNo: 123, OrderMoney: 22.1, OrderTime: &timestamppb.Timestamp{Seconds: time.Now().Unix()}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
