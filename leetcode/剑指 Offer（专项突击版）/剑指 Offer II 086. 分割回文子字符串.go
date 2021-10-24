package main

/*
给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。

647 回文子串，题目类似，但求的东西不一样，因而解法也不一样。（中心拓展）

这题用 回溯法
*/

func partition(s string) [][]string {
	res := [][]string{}
	n := len(s)
	var dfs func(int)
	path := []string{}
	dfs = func(left int) {
		if left == n {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := left; i < n; i++ {
			if check(left, i, s) {
				path = append(path, s[left:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func check(left, right int, s string) bool {
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
