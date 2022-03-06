package main

import "fmt"

func add(arr []int) {
	for i, _ := range arr {
		arr[i] += 5
	}
}

func main() {
	arr := make([]int, 0, 5)
	arr = append(arr, 1, 2, 3, 4, 5)
	add(arr)
	fmt.Print(arr)
}
