package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回 []。

排序+双指针
对数组进行排序。
遍历排序后数组：
	若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
	对于重复元素：跳过，避免出现重复解
	令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
		当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
		若和大于 0，说明 nums[R] 太大，R 左移
		若和小于 0，说明 nums[L] 太小，L 右移
*/

func threeSum(nums []int) [][]int {
	if len(nums) == 0 || len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := [][]int{}
	for i, one := range nums {
		if one > 0 {
			return res
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
