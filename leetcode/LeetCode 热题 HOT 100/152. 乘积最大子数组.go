package main

import "fmt"

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

当前位置如果是一个负数的话，那么我们希望以它前一个位置结尾的某个段的积也是个负数，这样就可以负负得正，并且我们希望这个积尽可能「负得更多」，即尽可能小。
如果当前位置是一个正数的话，我们更希望以它前一个位置结尾的某个段的积也是个正数，并且希望它尽可能地大。

*/
func maxProduct(nums []int) int {
	minF, maxF, ans := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		mi, ma := minF, maxF
		maxF = max(ma*num, max(num, mi*num))
		minF = min(mi*num, min(num, ma*num))
		ans = max(maxF, ans)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
