package main

import (
	"fmt"
	"math"
)

/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。
*/
func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	curSum := math.MinInt32
	for _, n := range nums {
		curSum = max(n, curSum+n)    //目前最大值的动态方程，之前的值是否要累加到当前值上
		maxSum = max(maxSum, curSum) //历史最大值
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
