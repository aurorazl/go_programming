package LeetCode_热题_HOT_100

import "math"

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
你可以认为每种硬币的数量是无限的。

问题的答案是通过子问题的最优解得到的。
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, one := range coins {
		for i := one; i <= amount; i++ {
			dp[i] = min(dp[i-one]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
