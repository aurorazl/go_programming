package main

import "fmt"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
动态规划
*/

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, i := range nums {
		for _, j := range res {
			tmp := make([]int, len(j))
			copy(tmp, j)
			tmp = append(tmp, i)
			res = append(res, tmp)
		}
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
