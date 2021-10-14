package main

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为O(n) 的算法解决此问题。
思路：
	利用空间换时间，引入字典
	两层循环，里层循环判断x是否是连续序列的第一个数，才进入。避免时间复杂度超过O（n）
*/

func longestConsecutive(nums []int) int {
	dict := map[int]bool{}
	for _, num := range nums {
		dict[num] = true
	}
	length := 0
	for k := range dict {
		if _, ok := dict[k-1]; !ok {
			curNum := k
			curLength := 1
			for dict[curNum+1] {
				curNum++
				curLength++
			}
			if curLength > length {
				length = curLength
			}
		}
	}
	return length
}
