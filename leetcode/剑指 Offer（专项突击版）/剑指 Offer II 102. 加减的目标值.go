package main

import "fmt"

/*
给定一个正整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

回溯法：
	backtrack(index+1, sum+nums[index])
	backtrack(index+1, sum-nums[index])

动态规划(0-1背包问题)
记数组的元素和为 sum，添加 - 号的元素之和为 neg，则其余添加 + 的元素之和为 sum−neg，得到的表达式的结果为
(sum−neg)−neg=sum−2⋅neg=target
neg = (sum-target)/2

定义二维数组 dp，其中 dp[i][j] 表示在数组 nums 的前 i 个数中选取元素，使得这些元素之和等于 j 的方案数。
假设数组 nums 的长度为 n，则最终答案为 dp[n][neg]。

dp[i][j]= dp[i−1][j]+dp[i−1][j−num]

*/

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1 // 边界条件，不选，一个解决方案
	for i := 0; i < n; i++ {
		for j := 0; j <= neg; j++ { // 从0开始，这样每行的i,0不用初始化
			dp[i+1][j] = dp[i][j]
			if nums[i] <= j {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[n][neg]
}

// 优化：由于 dp 的每一行的计算只和上一行有关，因此可以使用滚动数组的方式。内层循环需采用倒序遍历的方式

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
