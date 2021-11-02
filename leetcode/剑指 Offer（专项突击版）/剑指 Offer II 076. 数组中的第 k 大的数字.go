package main

import "fmt"

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

在长度为n的排序数组中，第k大的数字的下标是n-k
用快速排序的函数partition对数组分区，如果函数partition选取的中间值在分区之后的下标正好是n-k，分区后左边的的值都比中间值小，右边的值都比中间值大，
即使整个数组不是排序的，中间值也肯定是第k大的数字
如果函数partition选取的中间值在分区之后的下标大于n-k，那么第k大的数字一定位于中间值的左侧，于是再对中间值的左侧的子数组分区
如果函数partition选择的中间值在分区之后的下标小于n-k，那么第k大的数字一定位于中间值的右侧，于是再对中间值的右侧的子数组分区
*/

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	left, right := 0, len(nums)-1
	pos := partition(nums, left, right)
	for pos != n-k {
		if pos > n-k {
			right = pos - 1
		} else {
			left = pos + 1
		}
		pos = partition(nums, left, right)
	}
	return nums[pos]
}

func partition(nums []int, left int, right int) int {
	tmp := nums[left]
	for left < right {
		for left < right && nums[right] >= tmp {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= tmp {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = tmp
	return left
}
func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
