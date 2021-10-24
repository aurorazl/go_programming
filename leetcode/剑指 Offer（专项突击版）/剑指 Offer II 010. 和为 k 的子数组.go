package main

import "fmt"

/*
给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。

窗口滑动的条件是什么，while窗口内元素超过或者不满足条件时移动
但如果数组存在负数，遇到不满足题意的时候，我们应该移动窗口左边界，还是扩大窗口右边界从而寻找到符合条件的情况呢？
不确定

前缀和
当我们循环数组到下标N时，需要用到数组前N-1项的计算的结果（这里不一定非要是和，也可能是积等）
数组的前i个数字之和记为x，如果存在一个j(j<i)，数组的前j个数字之和为x-k，那么数组中从第i+1个数字开始到第j个数字结束的子数组之和为k
*/

func subarraySum(nums []int, k int) int {
	n := len(nums)
	cache := make(map[int]int)
	sum := 0
	res := 0
	cache[0] = 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		if cnt, ok := cache[sum-k]; ok {
			res += cnt
		}
		cache[sum] = cache[sum] + 1
	}
	return res
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}
