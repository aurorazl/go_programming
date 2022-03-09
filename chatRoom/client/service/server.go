package service

import (
	"fmt"
	"go_programming/chatRoom/common/utils"
	"net"
	"os"
)

/*
保持和server的通信
*/

func ShowMenu() {
	fmt.Println("-------恭喜xxx登录成功---------")
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")

	var key int

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表-")
	case 2:
		fmt.Println("你想对大家说的什么:)")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确..")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := utils.Transfer{Conn: conn}
	for {
		fmt.Println("client wait for server msg")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(mes)
	}

}
