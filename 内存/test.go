package main

import (
	"fmt"
	"unsafe"
)

type t1 struct {
	a [2]int8
	b int64
	c int16
}

type t2 struct {
	a [2]int8
	b int16
	c int64
}

func main() {
	fmt.Print(unsafe.Alignof(t1{}.c), unsafe.Sizeof(t1{}))
	fmt.Print(unsafe.Alignof(t2{}), unsafe.Sizeof(t2{}))
}
