package main

import "fmt"

func main() {
	var a = 10
	switch a {
	case 30:
		fmt.Println("ok")
		fallthrough
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 10:
		fmt.Println("ok3")
	}
	fmt.Println()
}
