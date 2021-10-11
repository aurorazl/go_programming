package main

import "fmt"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

1。暴力
对于每一个位置，我们都这样操作，得到一个矩形面积，求出它们的最大值。
	左边看一下，看最多能向左延伸多长，找到大于等于当前柱形高度的最左边元素的下标；
	右边看一下，看最多能向右延伸多长；找到大于等于当前柱形高度的最右边元素的下标。
接雨水不太一样，因为左边或者右边都是取最大。遍历一次即可，时间复杂度为O（n）

2。单调栈+哨兵
从左到右，下一个柱子高度严格小于上一个时，出栈，计算面积，
大于等于的时候入栈即可，计算最大面积会等于跳过中间的。
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

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}
