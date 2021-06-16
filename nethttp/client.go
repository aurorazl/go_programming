package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("收到服务端回复:%v\n", string(buf[:n]))
	}
}
