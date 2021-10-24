package main

import "fmt"

/*
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

前缀和
0表示为-1，其实就是求当前i的sum=map sum的最小index，是否更新maxLen
*/

func findMaxLength(nums []int) int {
	sum := 0
	cache := make(map[int]int)
	maxLen := 0
	cache[0] = -1
	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}
		if preIndex, has := cache[sum]; has {
			maxLen = max(maxLen, i-preIndex)
		} else { //	说明该这么多1的数量第一次出现
			cache[sum] = i
		}
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}
