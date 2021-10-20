package main

import "fmt"

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

二分查找（缺失数字以及后部分的值不等于index索引）（因为递增，不然可以考虑用数学公式）

*/

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
func missingNumber2(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return (len(nums)+1)*(len(nums))/2 - sum
}

func main() {
	fmt.Println(missingNumber2([]int{0, 1, 3, 4}))
}
