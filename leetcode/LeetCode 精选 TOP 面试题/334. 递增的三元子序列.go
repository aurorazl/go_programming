package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组nums ，判断这个数组中是否存在长度为 3 的递增子序列。
如果存在这样的三元组下标 (i, j, k)且满足 i < j < k ，使得nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
*/
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	small, mid := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= mid {
			mid = num
		} else if num > mid {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 3, 2, 4}))
}
