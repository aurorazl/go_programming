package main

import "fmt"

/*
一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
输入：m = 3, n = 7
输出：28
*/

//递归法
func uniquePathsRecurse(m int, n int) int {
	cnt := 0
	if m == 1 || n == 1 {
		return 1
	}
	if m > 1 {
		cnt += uniquePathsRecurse(m-1, n)
	}
	if n > 1 {
		cnt += uniquePathsRecurse(m, n-1)
	}
	return cnt
}

//非递归法
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
func main() {
	fmt.Println(uniquePaths(51, 9))
}
