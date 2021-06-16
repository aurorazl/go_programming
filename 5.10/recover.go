package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

	}()
	for {
		select {
		case <-ticker:
			proc()
		default:

		}
	}

}

func proc() {
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			fmt.Print(err)
		}
	}()
	panic("ok")
}
