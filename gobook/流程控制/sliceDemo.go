package main

import "fmt"

func main() {
	var a [2]int = [...]int{1, 2}
	var b = a
	fmt.Println(b)
}
