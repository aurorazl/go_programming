package main

import (
	"fmt"
	"math/rand"
)

func partition(li []int, start, end int) int {
	i := rand.Intn(end-start) + start
	li[start], li[i] = li[i], li[start]
	tmp := li[start]
	for start < end {
		for start < end && li[end] >= tmp {
			end--
		}
		li[start] = li[end]
		for start < end && li[start] <= tmp {
			start++
		}
		li[end] = li[start]
	}
	li[start] = tmp
	return start
}

func _quickSort(li []int, start, end int) {
	if start < end {
		mid := partition(li, start, end)
		_quickSort(li, start, mid-1)
		_quickSort(li, mid+1, end)
	}

}

func quickSort(li []int) {
	_quickSort(li, 0, len(li)-1)
}

func main() {
	li := []int{6, 1, 2, 4, 3}
	quickSort(li)
	fmt.Print(li)
}
