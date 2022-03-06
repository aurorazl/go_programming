package main

import (
	"fmt"
	"sort"
	"strings"
)

type a struct {
	name int
}

func ab() a {
	return a{}
}

type LIAN struct {
	from int
	to   int
}

func main() {

	var identifier []int = []int{1}
	var slice1 []int = make([]int, 1, 1)
	fmt.Print(identifier == nil, append(slice1, 1), slice1[1:], identifier[:])
	fmt.Println(&slice1[0])
	fmt.Println(strings.Contains("harbor.atlas.cn:8443/atlas/apulistech/tensorflow-serving@sha256:08a3a62d520eb4b96a9751ad07f07881477d3babf68211ae8e0a4b6d6dd77bea", "/"))
	fmt.Println(strings.SplitN("harbor.atlas.cn:8443/atlas/apulistech/tensorflow-serving@sha256:08a3a62d520eb4b96a9751ad07f07881477d3babf68211ae8e0a4b6d6dd77bea", "@", 2)[0])
	for _, one := range "ssss" {
		fmt.Println(one)
	}
	var b = make([]int, 2, 2)
	b[0] = 3
	c := &b[0]
	//e := b
	e := make([]int, 1)
	copy(e, b)
	b[1] = 2
	fmt.Println(&b[0])
	b = append(b, 1)
	fmt.Println(&b[0])
	fmt.Println(c, e)

	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	fmt.Println(b)

	a := make(map[string]interface{}, 1)
	a["s"] = struct{}{}
	d := (a)
	a["ss"] = struct{}{}
	fmt.Println(a["s"])
	fmt.Println(d)
}
