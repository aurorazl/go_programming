package main

import "fmt"

func main() {
	var a chan int
	a = make(chan int)
	go func() {
		a <- 1
	}()
	<-a
	var b []int
	var c []int
	fmt.Printf("%p %p", b, c)

}
