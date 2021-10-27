package main

import "math"

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

动态规划
上一步就只能在位置 (i - 1, j - 1) 或者位置 (i - 1, j)。

当 j=0 时 f[i][0]=f[i−1][0]+c[i][0]
f[i][j]=min(f[i−1][j−1],f[i−1][j])+c[i][j]
当 j=i 时 f[i][i]=f[i−1][i−1]+c[i][i]
*/

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j <= i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for _, num := range dp[m-1] {
		ans = min(ans, num)
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
