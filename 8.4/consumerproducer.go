package main

import (
	"fmt"
	"time"
)

var a = make(chan interface{}, 0)

func consumer() {

	for i := range a {
		fmt.Print(i)
	}

}

func producer() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		a <- i
	}
	close(a)
}

func main() {
	go producer()
	go consumer()
	time.Sleep(time.Second * 1)
}
