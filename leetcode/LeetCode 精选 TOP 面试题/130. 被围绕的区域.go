package main

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
被围绕的区间不会存在于边界上，换句话说，任何边界上的'O'都不会被填充为'X'。 任何不在边界上，或不与边界上的'O'相连的'O'最终都会被填充为'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

深度优先和广度优先
*/

func solve(board [][]byte) {
	if len(board) == 0 || len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs(board, 0, j)
		dfs(board, n-1, j)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[j][i] == 'A' {
				board[j][i] = 'O'
			} else if board[j][i] == 'O' {
				board[j][i] = 'X'
			}
		}
	}
}
func dfs(board [][]byte, x, y int) {
	if x >= 0 && y >= 0 && x < len(board[0]) && y < len(board) && board[y][x] == 'O' {
		board[y][x] = 'A'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)
	}
}
