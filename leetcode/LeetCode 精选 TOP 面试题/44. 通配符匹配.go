package main

import "fmt"

/*
给定一个字符串(s) 和一个字符模式(p) ，实现一个支持'?'和'*'的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意一个或多个字符串（包括空字符串）。

思路：
	如果 p.charAt(j) == s.charAt(i) : dp[i][j] = dp[i-1][j-1]；
	如果 p.charAt(j) == '？' : dp[i][j] = dp[i-1][j-1]；
	如果 p.charAt(j) == '*' :
		如果p.charAt(j-1) == s.charAt(i)，dp[i][j]=dp[i][j-1] || dp[i-1][j] .... 表示空字符串、或者aa
		如果p.charAt(j-1) != s.charAt(i)，dp[i][j]=dp[i-1][j]

	第三种情况也可以分为不使用星号dp[i][j]=dp[i][j-1]（空字符串），使用星号dp[i-1][j]
	// 为什么不使用dp[i-1][j-1]，因为j-1可能为空字符串，false，而i-1，j的边界已经处理过了

边界条件：
dp[0][0]=True，即当字符串 ss 和模式 pp 均为空时，匹配成功；
dp[i][0]=False，即空模式无法匹配非空字符串；
dp[0][j] 需要分情况讨论：因为星号才能匹配空字符串，所以只有当模式 p 的前 j 个字符均为星号时，dp[0][j] 才为真。
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Print(isMatch("abc", "*"))
}
