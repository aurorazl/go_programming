package main

import "fmt"

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	for i := heapSize/2 - 1; i > -1; i-- {
		buildMaxHeap(nums, i, heapSize-1)
	}
	for i := heapSize - 1; i > heapSize-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		buildMaxHeap(nums, 0, i-1)
	}
	return nums[0]
}
func buildMaxHeap(nums []int, low, high int) {
	tmp := nums[low]
	i := low
	j := 2*i + 1
	for j <= high {
		if j < high && nums[j] < nums[j+1] {
			j++
		}
		if nums[j] > tmp {
			nums[i] = nums[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	nums[i] = tmp
}
func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
