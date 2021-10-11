package main

/*
给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。
动态规划法
如果 p.charAt(j) == s.charAt(i) : dp[i][j] = dp[i-1][j-1]；
如果 p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1]；
如果 p.charAt(j) == '*'：
	如果 p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2] //就看j-2是否匹配了，j算是超过了i的长度，其他情况都不行了
	如果 p.charAt(j-1) == s.charAt(i) or p.charAt(j-1) == '.'：
		dp[i][j] = dp[i-1][j]    // a*可以代表aa，.*更强，可以一直归到[n][n+2]的情况（n是.*前面的元素个数，例如aabcd和aa.*）
		or dp[i][j] = dp[i][j-1] // a*可以代表一个a，.*更强，代表任意字符
		or dp[i][j] = dp[i][j-2] // a*表示空，.*更强，也可以代表空

*/
import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	dp := [][]bool{}
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(p)+1))
	}
	dp[0][0] = true
	dp[0][1] = false
	for j := 2; j < len(p)+1; j++ {
		if string(p[j-1]) == "*" {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			x := i - 1
			y := j - 1
			if p[y] == s[x] || string(p[y]) == "." {
				dp[i][j] = dp[i-1][j-1]

			} else if string(p[y]) == "*" {
				if p[y-1] != s[x] {
					dp[i][j] = dp[i][j-2]
				}
				if p[y-1] == s[x] || string(p[y-1]) == "." {
					dp[i][j] = dp[i][j-2] || dp[i-1][j] || dp[i][j-1]
				}
			}

		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	match, _ := regexp.MatchString("aa.*", "aa")
	fmt.Println(match)
	fmt.Println(isMatch("aa", "a*"))
}
