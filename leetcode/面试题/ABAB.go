package main

import (
	"fmt"
)

type test interface {
}

func main() {
	a := make(chan string)
	c := make(chan struct{})
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-a)
		}
		close(c)
	}()
	for i := 0; i < 10; i++ {
		a <- "A"
		a <- "B"
	}
	for range c {
		close(a)
	}

}
