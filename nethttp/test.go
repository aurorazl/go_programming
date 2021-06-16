package main

import "fmt"

type xxx interface {
	hello() int
}

func main() {
	var a = new([]int)
	var b bool
	var c = make([]bool, 0)
	var d []int
	//var e []int
	e := []int{1, 2, 3, 4}
	fmt.Print(*a == nil, b, c, d == nil)

	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		fmt.Println(v)
		s = append(s, v)
		fmt.Printf("len(s)=%v\n", len(s))
		fmt.Println(len(e))
	}
}
