package main

import (
	"fmt"
	. "grpcDemo/proto/package/hello"
	"io"
	"time"
)

type UserService struct {

}

func (*UserService) ReturnUsersByStream(in *UserRequest,stream UserService_ReturnUsersByStreamServer) error {
	var users []*UserInfo
	for index, user := range in.Users {
		users = append(users, user)
		if (index+1) % 2 == 0 {
			err := stream.Send(&UserResponse{Users: users})
			if err != nil {
				return err
			}
			users = users[0:0]
		}
		fmt.Println("sending")
		time.Sleep(time.Second*1)
	}
	if len(users)!=0 {
		stream.Send(&UserResponse{Users: users})
	}
	return nil
}

func (*UserService) SendUsersByStream(stream UserService_SendUsersByStreamServer) error {
	var users []*UserInfo
	for {
		userRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(users)
			return stream.SendAndClose(&UserResponse{Users: users})
		}
		if err != nil {
			return err
		}
		users = append(users, userRequest.Users...)
	}
}


func (*UserService) GetUsersByTWStream(stream UserService_GetUsersByTWStreamServer) error {
	var users []*UserInfo
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 这里不用处理，因为逻辑是直接将请求返回了
			return nil
		}
		if err != nil {
			fmt.Println("err is ",err)
			return err
		}
		users = append(users, req.Users...)
		err = stream.Send(&UserResponse{Users: users})
		users = users[0:0]
		if err != nil {
			fmt.Println("error is ",err)
			return err
		}
	}
}