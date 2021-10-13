package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	cur := 0
	start := 0
	length := math.MaxInt32
	left := 0
	fmt.Println(left)
	for i := 0; i < n; i++ {
		cur += nums[i]
		if cur >= target {
			for cur >= target {
				if i-start+1 < length {
					length = i - start + 1
					left = start
				}
				cur -= nums[start]
				start++
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
