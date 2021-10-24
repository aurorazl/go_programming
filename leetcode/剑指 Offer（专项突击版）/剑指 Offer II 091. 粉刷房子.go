package main

/*
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个n x 3的正整数矩阵 costs 来表示的。
例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2]表示第 1 号房子粉刷成绿色的花费，以此类推。
请计算出粉刷完所有房子最少的花费成本。

动态规划
dp[i][0] 代表 第i+1间房子涂红色时 前i+1间房子累计所需的 最小成本；
dp[i][1] 代表 第i+1间房子涂蓝色时 前i+1间房子累计所需的 最小成本；
dp[i][2] 代表 第i+1间房子涂绿色时 前i+1间房子累计所需的 最小成本；

dp[i][0] = Math.min(dp[i-1][1],dp[i-1][2])+costs[i][0];
dp[i][1] = Math.min(dp[i-1][0],dp[i-1][2])+costs[i][1];
dp[i][2] = Math.min(dp[i-1][0],dp[i-1][1])+costs[i][2];
*/

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][3]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i-1][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i-1][1]
		dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + costs[i-1][2]
	}
	return min(dp[n][0], min(dp[n][1], dp[n][2]))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
