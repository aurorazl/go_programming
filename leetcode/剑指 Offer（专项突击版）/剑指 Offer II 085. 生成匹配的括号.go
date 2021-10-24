package main

/*
正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

22.括号生成（回溯法）
*/

func generateParenthesis(n int) (ret []string) {
	var dfs func(l, r int, s string)
	dfs = func(l, r int, s string) {
		if n == r {
			ret = append(ret, s)
		}
		if l < n {
			dfs(l+1, r, s+"(") // 回溯，不用先增后删
		}
		if r < l {
			dfs(l, r+1, s+")")
		}
	}
	dfs(0, 0, "")
	return
}
