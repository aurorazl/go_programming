package main

import "fmt"

/*
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。

滑动窗口

类似 008.和大于等于target的最短子数组
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	ans := 0
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
		for left <= i && sum >= k { // 另一种做法，i为left，right不断尝试往右扩张。
			sum /= nums[left]
			left++
		}
		ans += i - left + 1 // ABCX，符合条件的有X，CX，BCX，ABCX共四个，个数符合right-left+1
	}
	return ans
}

// 另一种做法，i为left，right不断尝试往右扩张。
func numSubarrayProductLessThanK2(nums []int, k int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		right := i
		for sum < k {
			ans++
			if right < n-1 {
				right++
				sum *= nums[right]
			} else {
				break
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(numSubarrayProductLessThanK2([]int{10, 5, 2, 6}, 100))
}
