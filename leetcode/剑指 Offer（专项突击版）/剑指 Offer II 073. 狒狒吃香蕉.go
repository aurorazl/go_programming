package main

import (
	"fmt"
	"math"
)

/*

狒狒喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
狒狒可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉，下一个小时才会开始吃另一堆的香蕉。
狒狒喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

二分查找
k最大可能是piles中的最大值（h与piles.length相同时）
k最小可能是1
所以可以用二分查找，在1..max(piles)中选择一个值mid，
用piles[i]除以mid得到这堆香蕉能吃多长时间，一次猜测过程中需要合计所有的时间time
如果time<=h，时间有富裕，需要继续缩小mid的大小
如果time>h,时间不够，需要增加mid的大小
*/

func minEatingSpeed(piles []int, h int) int {
	start, end := 1, max(piles)
	for start < end {
		mid := (end + start) / 2
		if countTime(piles, mid) > h {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
func countTime(nums []int, k int) int {
	time := 0
	for _, num := range nums {
		time += (num + k - 1) / k
	}
	return time
}
func max(nums []int) int {
	ret := math.MinInt32
	for _, num := range nums {
		if num > ret {
			ret = num
		}
	}
	return ret
}
func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
