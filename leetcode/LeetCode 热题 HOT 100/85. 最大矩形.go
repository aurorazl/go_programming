package main

import "fmt"

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
*/

/*
解法，单调栈，每增加一行都是调用一次
*/
func largestRectangleArea(heights []int) int {
	stack := []int{0}
	tmp := make([]int, len(heights)+2)
	copy(tmp[1:len(heights)+1], heights)
	heights = tmp
	maxArea := 0
	for i := 1; i < len(heights); i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			curHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			maxArea = max(curHeight*(i-stack[len(stack)-1]-1), maxArea) //宽度算到上一个
		}
		stack = append(stack, i)
	}
	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([][]int, len(matrix))
	maxArea := 0
	heights[0] = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			heights[0][i] = 1
		} else {
			heights[0][i] = 0
		}
	}
	fmt.Println(heights)
	maxArea = max(largestRectangleArea(heights[0]), maxArea)
	for i := 1; i < len(matrix); i++ {
		heights[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[i][j] = 1 + heights[i-1][j]
			} else {
				heights[i][j] = 0
			}
		}
		maxArea = max(largestRectangleArea(heights[i]), maxArea)
	}
	return maxArea
}

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalRectangle([][]byte{{'1', '0'}}))
}
