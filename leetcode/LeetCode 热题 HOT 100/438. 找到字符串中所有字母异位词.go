package main

import "fmt"

/*
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指字母相同，但排列不同的字符串。

*/
func findAnagrams2(s string, p string) []int {
	needCnt := len(p)
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	ans := []int{}
	for i := 0; i < len(s); i++ {
		cur := i
		for i < len(s) && m[s[i]] != 0 {
			needCnt--
			m[s[i]]--
			i++
			if needCnt == 0 {
				ans = append(ans, i-len(p))
			}
		}
		m = make(map[byte]int)
		for j := 0; j < len(p); j++ {
			m[p[j]]++
		}
		needCnt = len(p)
		i = cur
	}
	return ans
}
func findAnagrams(s string, p string) []int {
	ans := []int{}
	if len(p) > len(s) {
		return ans
	}
	needCnt := len(p)
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	for i := 0; i < len(p); i++ {
		if m[s[i]] > 0 {
			needCnt--
		}
		m[s[i]]--
		if needCnt == 0 {
			ans = append(ans, i-len(p)+1)
		}
	}
	for i := len(p); i < len(s); i++ {
		if _, ok := m[s[i-len(p)]]; ok {
			m[s[i-len(p)]]++
			if m[s[i-len(p)]] > 0 {
				needCnt++
			}
		}
		if _, ok := m[s[i]]; ok {
			if m[s[i]] > 0 {
				needCnt--
			}
			m[s[i]]--
			if needCnt == 0 {
				ans = append(ans, i-len(p)+1)
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
