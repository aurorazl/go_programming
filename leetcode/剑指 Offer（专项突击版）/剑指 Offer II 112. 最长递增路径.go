package main

/*
给定一个m x n 整数矩阵matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

深度优先搜索，每个格子深入搜索

记忆化深度优先搜索，同一个单元格会被访问多次，每次访问都要重新计算。缓存。
由于同一个单元格对应的最长递增路径的长度是固定不变的，因此可以使用记忆化的方法进行优化。
*/
var (
	dirs = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	m, n int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n = len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(matrix, i, j, cache))
		}
	}
	return ans
}

func dfs(matrix [][]int, row int, column int, cache [][]int) int {
	if cache[row][column] != 0 {
		return cache[row][column]
	}
	cache[row][column]++ //本身是一个长度
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && matrix[newRow][newColumn] > matrix[row][column] {
			cache[row][column] = max(cache[row][column], dfs(matrix, newRow, newColumn, cache)+1) //遍历4个方向，深度搜索最大值后返回
		}
	}
	return cache[row][column]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
