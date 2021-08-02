package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

//双指针
func moveZeroes(nums []int) {
	slow, fast, n := 0, 0, len(nums)
	for fast < n {
		if nums[fast] == 0 {
			fast++ //右指针不断向右移动
		} else {
			nums[slow], nums[fast] = nums[fast], nums[slow] //每次右指针指向非零数，则将左右指针对应的数交换
			slow++
			fast++
		}
	}
}
func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
