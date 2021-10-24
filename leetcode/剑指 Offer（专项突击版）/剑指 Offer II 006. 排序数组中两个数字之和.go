package main

import "fmt"

/*
给定一个已按照 升序排列 的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。
函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。
numbers 的下标 从 0开始计数 ，所以答案数组应当满足 0<= answer[0] < answer[1] <numbers.length。
假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

方法一： 哈希
方法二： 遍历 + 二分查找另一个数
方法三： 双指针左右
*/
func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		index := binarySearch(numbers, target-num, i, len(numbers)-1)
		if index != -1 {
			return []int{i, index}
		}
	}
	return []int{}
}

func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 4, 9, 56, 90}, 8))
}
