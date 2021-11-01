package main

import "fmt"

/*
给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。

那就是如果mid所对应的一对数字下标是(奇数，偶数)，那么目标一定在mid之前，如果下标是(偶数，奇数)，目标一定在mid之后

*/

func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	n := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if mid > 0 && nums[mid] == nums[mid-1] {
			if mid%2 == 0 { // 说明在左边，忽略掉这一对
				right = mid - 2
			} else {
				left = mid + 1 //说明在右边
			}
		} else if mid < n-1 && nums[mid] == nums[mid+1] {
			if mid%2 == 0 {
				left = mid + 2
			} else {
				right = mid - 1
			}
		} else {
			return nums[mid]
		}
	}
	return nums[left]
}
func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))
}
