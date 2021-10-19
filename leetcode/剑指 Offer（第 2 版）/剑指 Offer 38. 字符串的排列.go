package main

import (
	"fmt"
	"sort"
)

/*
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

字符可能有重复，但输出的数组不能重复

动态规划,一层循环，动态方程
回溯法，递归遍历每一位，但要注意避免同一位放入重复的字符。
*/

func permutation(s string) []string {
	ret := [][]byte{}
	out := []string{}
	for index := range s {
		if len(ret) == 0 {
			ret = append(ret, []byte{s[index]})
		} else {
			tmp := [][]byte{}
			for i := range ret {
				one := ret[i]
				one_copy := make([]byte, len(one))
				for j := 0; j <= len(one); j++ {
					copy(one_copy, one) // 每次插入不同位置，copy防止前面的影响
					newOne := append(one_copy[:j], append([]byte{s[index]}, one_copy[j:]...)...)
					tmp = append(tmp, newOne)
				}
			}
			ret = tmp
		}
	}
	cache := make(map[string]bool, 0)
	for i := range ret {
		if !cache[string(ret[i])] {
			out = append(out, string(ret[i]))
			cache[string(ret[i])] = true
		}
	}
	return out
}

func permutation2(s string) []string {
	ans := []string{}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	n := len(s)
	perm := make([]byte, 0, n)
	vis := make([]bool, n) //对应每个字符是否访问过
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			ans = append(ans, string(perm)) // 遍历完一次
			return
		}
		for i, isVis := range vis { // 每次看一下每个字符没有被取过，重复也要考虑
			if isVis || i > 0 && t[i-1] == t[i] && !vis[i-1] {
				continue
			}
			perm = append(perm, t[i])
			vis[i] = true
			dfs(index + 1)
			perm = perm[:len(perm)-1] // perm len会不断变化
			vis[i] = false
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(permutation2("ab"))
}
