package main

import (
	"fmt"
	"sync"
	//"sync"
	//"time"
)

func test() {
	panic("11")
}

func main() {

	a := map[string]int{"11": 2, "112": 1}
	for k, _ := range a {
		fmt.Println(k)
		if k == "11" {
			//delete(a,"112")
			a["221"] = 1
		}
	}
	fmt.Print(a)

	b := []int{1, 2, 3}
	for _, i := range b {
		fmt.Print(i)
		b[1] = 0
	}
	fmt.Print(len(a))

	var c sync.Map
	c.Store(1, "a")
	//d :=c
	//fmt.Print(d)

	//var e sync.Mutex
	//go func() {
	//	defer func() {
	//		if p := recover();p!=nil{
	//			fmt.Print(p)
	//		}
	//	}()
	//	e.Lock()
	//	test()
	//	e.Unlock()
	//}()
	//time.Sleep(time.Second)
	//e.Lock()
	//fmt.Print()
	//e.Unlock()

	f := make(chan int, 1)
	close(f)
	close(f)
}
