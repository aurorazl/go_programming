package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

本质上是一个排序问题
若拼接字符串 x + y > y + x，则 x “大于” y ；
传递性证明
*/

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
		b, _ := strconv.Atoi(strconv.Itoa(nums[j]) + strconv.Itoa(nums[i]))
		return a < b
	})
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), ""), "[]")
}

func main() {
	fmt.Println(minNumber([]int{3, 30}))
}
