package main

import (
	"fmt"
)

//一次性锁
type Mutex chan struct{}

func newMutex() Mutex {
	return make(Mutex, 1)
}

var o = newMutex()

func (o Mutex) Lock() {
	o <- struct{}{}
}

func (o Mutex) UnLock() {
	<-o
}

//

// Once

type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)
	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	_, ok := <-o
	if !ok {
		return
	}
	f()
	close(o)
}

//读写锁
type RWMutex struct {
	write chan struct{}
	read  chan int
}

func NewRwmutex() RWMutex {
	return RWMutex{
		write: make(chan struct{}, 1),
		read:  make(chan int, 1),
	}
}

func (rw RWMutex) Lock()   { <-rw.write }
func (rw RWMutex) Unlock() { rw.write <- struct{}{} }

func (rw RWMutex) Rlock() {
	var rs int
	select {
	case rw.write <- struct{}{}:
	case rs = <-rw.read:
	}
	rs++
	rw.read <- rs
}
func (rw RWMutex) Runlock() {
	rs := <-rw.read
	rs--
	if rs == 0 {
		<-rw.write
		return
	}
	rw.read <- rs
}

// WaitGroup
type WaitGroup struct {
	cnt chan int
	end chan struct{}
}

func NewWaitGroup() WaitGroup {
	wg := WaitGroup{
		cnt: make(chan int, 1),
		end: make(chan struct{}),
	}
	wg.cnt <- 0
	return wg
}

func (w WaitGroup) Add(i int) {
	var rs int
	rs = <-w.cnt
	rs += i
	if rs == 0 {
		close(w.end)
	}
	if rs < 0 {
		panic("<0")
	}
	w.cnt <- rs
}

func (w WaitGroup) Done() { w.Add(-1) }
func (w WaitGroup) Wait() {
	fmt.Println("waiting")
	<-w.end
}

func main() {
	var a chan int
	a = make(chan int, 1)
	a <- 1
	defer func() {
		fmt.Println("closing")
		close(a)
	}()
	fmt.Print("hello")

	var b = NewWaitGroup()
	for i := 0; i < 3; i++ {
		b.Add(1)
		go func(no int) {
			fmt.Println("sub groutine ", no)
			b.Done()
		}(i)
	}

	b.Wait()

	select {
	default:
		fmt.Println("case 1")
	}

}
