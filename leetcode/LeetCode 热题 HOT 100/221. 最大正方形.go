package main

import (
	"fmt"
	"math"
)

/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

动态规划
如果该位置的值是 00，则 dp(i,j)=0，因为当前位置不可能在由 11 组成的正方形中；
如果该位置的值是 11，则 dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 11，状态转移方程如下：
	dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

*/
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	maxSide := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(li ...int) int {
	minV := math.MaxInt32
	for _, i := range li {
		if i < minV {
			minV = i
		}
	}
	return minV
}
func main() {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
