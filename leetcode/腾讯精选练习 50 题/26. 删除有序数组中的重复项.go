package main

import "fmt"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

双指针
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	j := 1
	for ; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 1}))
}
