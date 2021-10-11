package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
解析
	下雨后水能到达的最大高度等于下标 i 两边的最大高度的最小值，下标 i 处能接的雨水量等于下标 i 处的水能到达的最大高度减去 height[i]。
	两边的最大高度分别扫两次扫出来
再遍历一次即可
创建两个长度为 nn 的数组 leftMax 和 rightMax。对于0≤i<n，leftMax[i] 表示下标 i 及其左边的位置中，height 的最大高度，
rightMax[i] 表示下标 i 及其右边的位置中，height 的最大高度。
显然，leftMax[0]=height[0]，rightMax[n−1]=height[n−1]。两个数组的其余元素的计算如下：
	当 1≤i≤n−1 时，leftMax[i]=max(leftMax[i−1],height[i])；
	当 0≤i≤n−2 时，rightMax[i]=max(rightMax[i+1],height[i])。

*/

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	storage := make([]int, len(height))
	for index, _ := range height {
		h := min(leftMax[index], rightMax[index]) - height[index]
		if h > 0 {
			storage[index] = h
		}
	}
	return sum(storage)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func sum(list []int) int {
	s := 0
	for _, i := range list {
		s += i
	}
	return s
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
