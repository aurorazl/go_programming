package main

/*
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。

动态规划
回溯+数组vis标记
*/

func permute(nums []int) (res [][]int) {
	n := len(nums)
	vis := make([]bool, n)
	perm := []int{}
	var dfs func(int)
	dfs = func(cnt int) {
		if cnt == n {
			res = append(res, append([]int{}, perm...))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			perm = append(perm, nums[i])
			dfs(cnt + 1)
			perm = perm[:len(perm)-1]
			vis[i] = false
		}
	}
	dfs(0)
	return
}
