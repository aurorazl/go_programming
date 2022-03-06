package main

import (
	"fmt"
)

func test(s []int) {
	s = append(s, 4)
	s[0] = 0
}

func main() {
	var b []int
	b = append(b, 1)
	a := make([]int, 0)
	fmt.Print(a, b)

	s := []int{1, 2, 3}
	test(s)
	fmt.Println(s)

	c := make([]int, 10)
	d := c[:2]
	d = append(d, 2)
	fmt.Println(d, c)

}
