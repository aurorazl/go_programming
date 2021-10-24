package main

/*
给你一个整数数组nums ，请计算数组的 中心下标 。
数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。

前缀和，循环直到最后一个
2*sum+nums[i] == total
*/

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if 2*sum+nums[i] == total { // 也能判断空和的情况，因为一开始先判断，再加
			return i
		}
		sum += nums[i]
	}
	return -1
}
