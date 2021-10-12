package main

import "fmt"

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
输入：nums = [3,2,1,0,4]
输出：false
*/

//递归法
func canJumpReverse(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for i := 1; i <= nums[0]; i++ {
		if canJumpReverse(nums[i:]) {
			return true
		}
	}
	return false
}

//非递归
func canJump(nums []int) bool {
	rightMax := 0
	for index, i := range nums {
		fmt.Println(index, i, rightMax)
		if index <= rightMax {
			rightMax = max(rightMax, index+i)
			if rightMax >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
