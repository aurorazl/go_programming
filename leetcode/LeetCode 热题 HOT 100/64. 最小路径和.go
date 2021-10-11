package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
动态规划
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(grid[0]))
		if i == 0 {
			dp[0][0] = grid[0][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	fmt.Println(dp)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
