package main

import "fmt"

/*
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

// 方法一：中心扩散法。
func longestPalindrome(s string) string {
	start := 0
	maxLength := 1
	for i := 1; i < len(s); i++ {
		left := i - 1
		right := i + 1
		length := 1
		for left >= 0 && s[left] == s[i] {
			left--
			length++
		}
		for right < len(s) && s[right] == s[i] {
			right++
			length++
		}
		for s[left] == s[right] {
			left--
			right++
			length += 2
		}
		if maxLength < length {
			maxLength = length
			start = left
		}
	}
	return s[start : start+maxLength]
}

func main() {
	fmt.Println(longestPalindrome("bab"))
}
