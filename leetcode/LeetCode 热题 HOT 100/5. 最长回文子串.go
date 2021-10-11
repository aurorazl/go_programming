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
	maxLength := 0
	for i := 0; i < len(s); i++ {
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
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length += 2
			left--
			right++
		}
		if maxLength < length {
			maxLength = length
			start = left
		}
	}
	return s[start+1 : start+maxLength+1]
}

//方法2： 动态规划
func longestPalindrome2(s string) string {
	dp := [][]bool{}
	for _ = range s {
		dp = append(dp, make([]bool, len(s)))
	}
	maxLen := 1
	start := 0
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[i-1][j+1]) == true {
				dp[i][j] = true
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					start = j
				}
			}
		}
	}
	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindrome2("abbc"))
}
