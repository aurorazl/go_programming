package main

import "fmt"

/*
现有⼀个数组 unsigned v[]，v[i] 表示最⼤移动距离，⽐如 v[1]=2表示可以从位置1⾛
到2和3，但⾛不到4。
要求：写⼀个函数，判断从位置0开始，能否⾛到最后⼀个位置？
*/

func CanReach(arr []int, size int) bool {
	if len(arr) == 0 || (len(arr) == 1 && arr[0] >= 1) {
		return true
	}
	curNum := arr[0]
	if curNum == 0 {
		return false
	}
	for i := 1; i <= curNum; i++ {
		if CanReach(arr[i:], len(arr)-i) {
			return true
		}
	}
	return false
}

func main() {
	a := []int{2, 10, 1, 1, 4}
	b := []int{2, 2, 1, 0, 4}
	fmt.Print(CanReach(a, len(a)))
	fmt.Print(CanReach(b, len(b)))
}
