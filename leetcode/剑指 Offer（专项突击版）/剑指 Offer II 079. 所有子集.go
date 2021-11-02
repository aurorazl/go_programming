package main

/*
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

回溯
*/

func subsets(nums []int) [][]int {
	ret := [][]int{}
	var dfs func(x int)
	path := []int{}
	n := len(nums)
	dfs = func(x int) {
		if x == n {
			ret = append(ret, append([]int{}, path...))
			return
		}
		path = append(path, nums[x])
		dfs(x + 1)
		path = path[:len(path)-1]
		dfs(x + 1)

	}
	dfs(0)
	return ret
}
