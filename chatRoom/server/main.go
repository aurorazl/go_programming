package main

import (
	"fmt"
	"go_programming/chatRoom/server/controller"
	"go_programming/chatRoom/server/dao"
	"go_programming/chatRoom/server/database"
	"log"
	"net"
)

func initDao() {
	dao.MyUserDao = dao.NewUserDao(database.RedisConn, database.Ctx)
}

func main() {
	initDao()
	fmt.Println("start listening on 8899")
	listen, err := net.Listen("tcp", "0.0.0.0:8899")
	defer listen.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		processor := controller.Processor{Conn: accept}
		go processor.Process()
	}
}
