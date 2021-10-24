package main

import "fmt"

/*
给定两个字符串s和p，找到s中所有 p 的变位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
变位词 指字母相同，但排列不同的字符串。

14题的拓展，滑动窗口，满足条件加入结果中
*/
func findAnagrams(s string, p string) (res []int) {
	m, n := len(p), len(s)
	if m > n {
		return
	}
	var cnt [26]int // 标记每个字符的数量
	for i, char := range p {
		cnt[char-'a']--
		cnt[s[i]-'a']++ // s2初始化前n个
	}
	diff := 0
	for _, count := range cnt {
		if count != 0 {
			diff++
		}
	}
	if diff == 0 {
		res = append(res, 0)
	}
	for i := m; i < n; i++ {
		x, y := s[i]-'a', s[i-m]-'a'
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			res = append(res, i-m+1)
		}
	}
	return
}
func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
