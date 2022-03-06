package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"runtime/pprof"
	"sync"
	"time"
)

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	ch := make(chan struct{})
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func(n int) {
			mutex.Lock()
			fmt.Println(n, windows.GetCurrentThreadId())
			mutex.Unlock()
			//fmt.Printf(("threads in starting: %d\n"), threadProfile.Count())
			time.Sleep(time.Second * 5)
			mutex.Lock()
			fmt.Println(n, windows.GetCurrentThreadId())
			mutex.Unlock()
			//fmt.Printf(("threads in starting: %d\n"), threadProfile.Count())
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Print(pprof.Label)

}
