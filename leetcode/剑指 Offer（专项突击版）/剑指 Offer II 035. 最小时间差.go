package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

一天有24小时，1440分钟，开一个1440长度的布尔数组模拟哈希表，把时间换算成0~1439之间的数值，将数值对应数组中的位置设置为true；
遍历数组，找离得最近的两个时间点。

*/

func findMinDifference(timePoints []string) int {
	array := [1440]int{}
	for _, point := range timePoints {
		splits := strings.Split(point, ":")
		hour, _ := strconv.Atoi(splits[0])
		minutes, _ := strconv.Atoi(splits[1])
		value := hour*60 + minutes
		if array[value] != 0 {
			return 0
		}
		array[value] = 1
	}
	minLen := math.MaxInt32
	prev := -1
	first := -1
	last := -1
	for i := 0; i < 1440; i++ {
		if array[i] != 0 {
			if prev == -1 {
				prev = i
				first = i
			} else {
				minLen = min(minLen, i-prev)
				prev = i
				last = i
			}
		}
	}
	return min(minLen, 1440-last+first) // 首尾两个值可能差距更小
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}
