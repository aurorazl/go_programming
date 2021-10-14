package main

import "fmt"

/*
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定s 是否可以被空格拆分为一个或多个在字典中出现的单词。
说明：
	拆分时可以重复使用字典中的单词。
	你可以假设字典中没有重复的单词。
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
思路：
	一个字符串前后顺序看看是否在里面，动态规划
*/
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if dp[i] && dict[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
