package main

import "fmt"

var a = b + c // a 第三个初始化, 为 3
var b = 2
var c = 1 // c 第一个初始化, 为 1

func init() {
	var a map[string]int
	b, _ = a[""], true
}

func main() {
	var a = `s`
	fmt.Print(b, a, (int)(0))
}
