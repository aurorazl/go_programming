package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target(可能重复)。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
进阶：
	你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

二分
二分查找中，寻找 leftIdx 即为在数组中寻找第一个大于等于 target 的下标，
寻找 rightIdx 即为在数组中寻找第一个大于 target 的下标，然后将下标减一。

*/
func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)
	if left <= right && right < len(nums) && nums[left] == nums[right] {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left := 0
	right := len(nums) - 1
	ans := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target { //特殊处理相等的情况
			ans = mid
			if lower {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
