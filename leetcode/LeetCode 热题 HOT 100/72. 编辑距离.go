package main

import "fmt"

/*
给你两个单词word1 和word2，请你计算出将word1转换成word2 所使用的最少操作数。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
	horse -> rorse (将 'h' 替换为 'r')
	rorse -> rose (删除 'r')
	rose -> ros (删除 'e')
思路：
A和B单词的编辑距离等于
	在单词 A 中插入一个字符；i-1 j
	在单词 B 中插入一个字符； j-1 i
	修改单词 A 的一个字符。 i-1 j-1 看增加的单词是否相同，不相同+1
*/

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]-1)) + 1
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	fmt.Println(minDistance("horse", "ros"))
}
