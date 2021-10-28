package main

import "fmt"

/*
给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。

解法：
动态规划

能否选择若干物品，使它们刚好放满一个容量为 t 的背包。若每种物品只有一个，则为 0 - 1 背包问题；
若每个物品的个数有限，则为多重背包问题；
若每个物品的个数无限，则为完全背包问题。

dp[i][j] = dp[i-1][j] 前面的元素已经满足 || dp[i-1][j-nums[i-1]] 用上当前元素
第一层循环是nums，这样不用去重
第二层循环是target

*/
func canPartition(nums []int) bool {
	n := len(nums)
	dp := make([][]bool, n+1)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j] || (j-nums[i-1] >= 0 && dp[i-1][j-nums[i-1]])
		}
	}
	return dp[n][target]
}

// 优化：当前行其实是在前一行的基础上进行更新的，所以使用一维的数组可以无需复制前一行的数据直接更新，这样会更高效。但是要注意 j 是从大往小遍历，因为这样不会破坏之前的值。
func canPartition2(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := target; j >= 1; j-- {
			dp[j] = dp[j] ||
				(j-nums[i-1] >= 0 && dp[j-nums[i-1]])
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(canPartition2([]int{3, 2, 1, 6}))
}
