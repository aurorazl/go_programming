package main

import "fmt"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

*/
func maxSubArray(nums []int) int {
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println(i, nums[i], nums[i-1])
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}
	return maxValue
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
