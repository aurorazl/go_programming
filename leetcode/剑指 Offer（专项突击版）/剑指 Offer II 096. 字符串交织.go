package main

import "fmt"

/*
给定三个字符串s1、s2、s3，请判断s3能不能由s1和s2交织（交错）组成。
两个字符串 s 和 t 交织的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
提示：a + b 意味着字符串 a 和 b 连接。


回溯法，尝试每种可能，直到s3尾部
动态规划
s3的第i个字符，都有可能从s1或者s2来（不用确定，只需要直到第i个字符能够实现即可）
动态发程：
	dp[i + 1][j + 1] = ((ch1 == ch3) && dp[i][j + 1]) || ((ch2 == ch3) && dp[i + 1][j]);

*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	if m == 0 || n == 0 {
		return s1+s2 == s3
	}
	dp := make([][]bool, m+1) // 表示当前用了几个s1，几个s2
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m && s1[i] == s3[i]; i++ {
		dp[i+1][0] = true
	}
	for j := 0; j < n && s2[j] == s3[j]; j++ {
		dp[0][j+1] = true
	}
	for i := 1; i <= m; i++ { // 用了几个单词
		for j := 1; j <= n; j++ {
			ch1 := s1[i-1]
			ch2 := s2[j-1]
			ch3 := s3[i+j-1]
			dp[i][j] = (ch1 == ch3 && dp[i-1][j]) || (ch2 == ch3 && dp[i][j-1])
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(isInterleave("db", "b", "cbb"))
}
