package main

import "sync"

func makeThumbnails6(filenames <-chan string) int {
	sizes := make(chan int)
	var wg sync.WaitGroup // 用channel取n次值也可以
	i := 0
	for _ = range filenames {
		wg.Add(1)
		i++
		go func() {
			defer wg.Done()
			sizes <- i
		}()

	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int
	for size := range sizes {
		total += size
	}
	return total
}

// 自己实现
