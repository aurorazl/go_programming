package main

import "fmt"

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

快慢双指针
左右双指针
*/

//快慢
func exchange(nums []int) []int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j]%2 == 0 {
			j++
		} else {
			if i == j {
				j++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
