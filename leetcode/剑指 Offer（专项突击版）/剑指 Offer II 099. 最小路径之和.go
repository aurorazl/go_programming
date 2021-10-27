package main

import "fmt"

/*
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：一个机器人每次只能向下或者向右移动一步。

动态规划
（一般动态规划的问题都可以通过穷举的方式得到答案，既然可以穷举，我们就可以将这个问题抽象成树形结构问题，然后通过回溯来解决。只是回溯时间复杂度更高）

dp[i][j] = grid[i][j] + Math.min(dp[i + 1][j], dp[i][j + 1]);
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[0][0] = grid[0][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
