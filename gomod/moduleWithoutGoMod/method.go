package main

import (
	"fmt"
	"moduleWithoutGoMod/DirectoryA"
)

func Test() {
	fmt.Println("module without go mod")

}

func main() {
	DirectoryA.Test()
}