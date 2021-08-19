package main

import "fmt"

/*
给定一个包含n + 1 个整数的数组nums ，其数字都在 1 到 n之间（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。
你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。
nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

快慢指针，数组的数字表示下一次的位置，由于至少有两个相同的数，他们指向同一个位置，肯定存在环
*/
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for ; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {

	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

// 每个数都可以对应数组的一个下标。
func findDuplicate2(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		} else {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 2}))
	fmt.Println(findDuplicate2([]int{1, 2, 2}))
}
