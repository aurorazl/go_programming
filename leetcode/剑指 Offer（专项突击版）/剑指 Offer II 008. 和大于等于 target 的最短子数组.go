package main

/*
给定一个含有n个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0

窗口内之和小于目标值时右指针向右移动增加窗口宽度；
窗口内之和大于目标值时左指针向右移动减少窗口宽度。
*/
import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	cur := 0
	start := 0
	length := math.MaxInt32
	for i := 0; i < n; i++ {
		cur += nums[i]
		if cur >= target { //满足情况了
			for cur >= target { //有可能左指针右移仍然满足情况。
				if i-start+1 < length {
					length = i - start + 1
				}
				cur -= nums[start]
				start++ //左指针右移
			}
		}
	}
	if length == math.MaxInt32 {
		return 0
	}
	return length
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 4, 3}))
}
