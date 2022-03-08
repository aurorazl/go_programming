package main

import "fmt"

type A struct {
	name string
	age  int
}
type B []int

func (b B) test() {
	b[0] = 2
}

func main() {
	//var a [2]int = [...]int{1, 2}
	//var b = A{"guanyu",23}
	var b B = []int{1, 2}
	fmt.Println(b)
	b.test()
	fmt.Println(b)
}
