package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的 最短 子数组，并输出它的长度。
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

其实就是要找“断崖”
*/
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	MA, MI := math.MinInt32, math.MaxInt32
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if MA > nums[i] {
			l = i
		} else {
			MA = nums[i] //开始和结束，l的位置都不会在变化，中间会变化，那么找到的是右边界
		}
	}
	for i := n - 1; i >= 0; i-- {
		if MI < nums[i] {
			r = i
		} else {
			MI = nums[i] //开始和结束，r的位置都不会在变化，中间会变化，那么找到的是右边界
		}
	}
	if r == l {
		return 0
	}
	return l - r + 1
}
func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
