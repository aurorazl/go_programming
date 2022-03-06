package main

import "fmt"

func main() {
label:
	var a = 1
	fmt.Println(a)
	goto label
}
