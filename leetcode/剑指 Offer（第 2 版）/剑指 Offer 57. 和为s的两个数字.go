package main

/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

双指针法
*/

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}
