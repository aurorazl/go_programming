package main

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

回溯
*/

func combine(n int, k int) [][]int {
	ret := [][]int{}
	var dfs func(x int)
	path := []int{}
	dfs = func(x int) {
		if len(path) == k {
			ret = append(ret, append([]int{}, path...))
			return
		}
		if x == n+1 {
			return
		}
		path = append(path, x)
		dfs(x + 1)
		path = path[:len(path)-1]
		dfs(x + 1)

	}
	dfs(1)
	return ret
}
