package main

import (
	"fmt"
)

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

思路：
	遍历找断崖
	二分查找
*/

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}

func minArray2(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := (left + right) / 2 // 由于可能有余数，mid偏小
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right-- // 两边都有可能，那对于两种情况，左边都可以再收缩一次。
		}
	}
	return numbers[left]
}
func main() {
	fmt.Println(minArray2([]int{2, 2, 2, 0, 1}))
}
