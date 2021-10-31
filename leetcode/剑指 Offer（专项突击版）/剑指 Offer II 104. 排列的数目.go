package main

import "fmt"

/*
给定一个由 不同正整数组成的数组 nums ，和一个目标整数 target 。
请从 nums 中找出并返回总和为 target 的元素组合的个数。数组中的数字可以在一次排列中出现任意次，但是顺序不同的序列被视作不同的组合。

动态规划（完全背包问题）

与最少的硬币数量解法类似，但要求出全部排列，所以coin必须在里层

*/

func combinationSum4_erro(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums { // 这种解法是错误的，因为nums的出现是有顺序的。
		for j := num; j <= target; j++ {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for _, num := range nums {
			if num <= j {
				dp[j] += dp[j-num]
			}
		}
	}
	return dp[target]
}
func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
