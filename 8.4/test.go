package main

import "fmt"

var ch = make(chan string, 1)

func main() {
	fmt.Println(cap(ch))
	ch <- "A"
	fmt.Println(len(ch))
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(i, x) // "0" "2" "4" "6" "8"
		case ch <- i:
			fmt.Println(111, i)
		}
	}
}
