package main

import "fmt"

/*
给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在O(n) 时间复杂度内完成此题。(不能先完全乘，然后再除)
进阶：
	你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
1。
利用索引左侧所有数字的乘积和右侧所有数字的乘积（即前缀与后缀）相乘得到答案。
3个数组
2。
将 L 或 R 数组用输出数组来计算。先把输出数组当作 L 数组来计算，然后再动态构造 R 数组得到结果。
*/

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}
	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R = R * nums[i]
	}
	return ans
}

func main() {
	fmt.Print(productExceptSelf([]int{1, 2, 3, 4}))
}
