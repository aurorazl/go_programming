package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。

滑动窗口
遍历右指针，左指针移动直到新增字符为0
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	left := 0
	right := 0
	maxLen := 0
	cache := make(map[byte]bool)
	for right < n {
		char := s[right]
		for cache[char] {
			cache[s[left]] = false
			left++
		}
		cache[char] = true
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
