package main

import "fmt"

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

回溯思想
*/

func exist(board [][]byte, word string) bool {
	mark := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		mark[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				mark[i][j] = 1
				if track(board, i, j, word[1:], mark) {
					return true
				}
				mark[i][j] = 0
			}
		}
	}
	return false
}

func track(board [][]byte, i int, j int, word string, mark [][]int) bool {
	if len(word) == 0 {
		return true
	}
	//上下左右
	for _, one := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		next_i := one[0] + i
		next_j := one[1] + j
		if next_i >= 0 && next_j < len(board[0]) && next_i < len(board) && next_j >= 0 && board[next_i][next_j] == word[0] {
			if mark[next_i][next_j] == 1 {
				continue
			}
			mark[next_i][next_j] = 1
			if track(board, next_i, next_j, word[1:], mark) {
				return true
			}
			mark[next_i][next_j] = 0
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
