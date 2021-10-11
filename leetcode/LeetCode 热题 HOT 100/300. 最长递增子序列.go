package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

当 nums[i] > nums[j] 时： nums[i]可以接在 nums[j] 之后（此题要求严格递增），此情况下最长上升子序列长度为 dp[j] + 1 ；
当 nums[i] <= nums[j] 时： nums[i] 无法接在nums[j] 之后，此情况上升子序列不成立，跳过。

转移方程： dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)。看看接在哪个上

*/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return maxList(dp...)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxList(li ...int) int {
	maxV := math.MinInt32
	for _, i := range li {
		if i > maxV {
			maxV = i
		}
	}
	return maxV
}
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
