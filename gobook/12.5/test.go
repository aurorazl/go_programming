package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("fd")
	fmt.Println(fd.Int()) // "1"
	fd.SetInt(2)
}
