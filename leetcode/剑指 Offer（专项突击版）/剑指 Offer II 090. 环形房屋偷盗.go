package main

import "fmt"

/*
一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

与	剑指 Offer II 089. 房屋偷盗 	比较，多了限制条件，第一家和最后一家不能同时偷

第一家偷的最高金额为dynamic(nums, 0, n - 1)
第一家不偷的最高金额为dynamic(nums, 1, n)
*/

func rob(nums []int) int {
	n := len(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(dynamic(nums[:n-1]), dynamic(nums[1:n]))
}

func dynamic(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{0}))
}
