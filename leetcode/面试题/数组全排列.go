package main

import "fmt"

func findAllList(list []int) [][]int {
	storage := [][]int{}
	for _, i := range list {
		if len(storage) == 0 {
			storage = append(storage, []int{i})
		} else {
			tmp := [][]int{}
			for _, j := range storage {
				j_copy := make([]int, len(j))
				for index := 0; index <= len(j); index++ {
					copy(j_copy, j)
					tmp = append(tmp, append(j_copy[0:index], append([]int{i}, j_copy[index:]...)...))
				}
			}
			storage = tmp
		}
	}
	return storage
}

func main() {
	list := findAllList([]int{1, 2, 3, 4})
	fmt.Println(list)
}
