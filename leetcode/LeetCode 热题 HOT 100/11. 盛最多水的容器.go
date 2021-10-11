package main

import "fmt"

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
在每一个状态下，无论长板或短板收窄 1 格，都会导致水槽 底边宽度−1：
	若向内移动短板，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此水槽面积 S(i, j)S(i,j) 可能增大。
	若向内移动长板，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 不变或变小，下个水槽的面积一定小于当前水槽面积。
*/

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max_area := 0
	for i < j {
		if height[i] < height[j] {
			max_area = max2(height[i]*(j-i), max_area)
			i++
		} else {
			max_area = max2(height[j]*(j-i), max_area)
			j--
		}
	}
	return max_area
}

func max2(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
