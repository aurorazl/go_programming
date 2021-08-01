package main

import (
	"fmt"
	"math"
)

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为O(n) 的算法解决此问题。
*/

func longestConsecutive(nums []int) int {
	storage := make(map[int]bool)
	maxLen := math.MinInt32
	for _, num := range nums {
		storage[num] = true
	}
	for num := range storage {
		if !storage[num-1] { // 只有当一个数是连续序列的第一个数的情况下才会进入内层循环,然后在内层循环中匹配连续序列中的数
			curNum := num
			curLen := 1
			for storage[curNum+1] {
				curNum++
				curLen++
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
