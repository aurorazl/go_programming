package main

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

动态规划
f(i,j)=max[f(i,j−1),f(i−1,j)]+grid(i,j)
*/

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0] // 初始化，减小判断次数
	}
	for j := 1; j < n; j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = max(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
