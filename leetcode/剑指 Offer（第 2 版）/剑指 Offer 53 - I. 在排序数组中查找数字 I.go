package main

import "fmt"

/*
统计一个数字在排序数组中出现的次数。

二分查找
要找的就是数组中「第一个等于 target 的位置」（记为leftIdx）和「第一个大于 target 的位置减一」（记为 rightIdx）。
寻找 leftIdx 即为在数组中寻找第一个大于等于 target 的下标，寻找 rightIdx 即为在数组中寻找第一个大于 target 的下标，然后将下标减一。
--> left index 查找target的最左位置，right index查找target+1的最左位置即可。
二分查找两个<=target的最左边界
*/

func search(nums []int, target int) int {
	return helper(nums, target+1) - helper(nums, target)
}

// 一个函数，查找target最左边的位置
func helper(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	fmt.Println(search([]int{1, 2, 2, 4}, 3))
}
