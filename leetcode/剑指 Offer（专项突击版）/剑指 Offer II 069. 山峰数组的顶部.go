package main

import "fmt"

/*
符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：

arr.length >= 3
存在 i（0 < i< arr.length - 1）使得：
	arr[0] < arr[1] < ... arr[i-1] < arr[i]
	arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i，即山峰顶部。

枚举，直到arr[i] < arr[i + 1]

*/

func peakIndexInMountainArray(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	fmt.Println(peakIndexInMountainArray([]int{1, 3, 5, 4, 2}))
}
