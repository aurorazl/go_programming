package main

import (
	"fmt"
	"math"
)

/*
给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的 最少分割次数 。

方法一： 86 分割回文子字符串  回溯法，所有解法中长度最小的。（会超出时间限制）
方法二： 动态规划
	循环l,r遍历所有情况，dp[r] = min(dp[r],dp[l]+1)

*/
func minCut2(s string) int {
	n := len(s)
	ret := math.MaxInt32
	var dfs func(int)
	path := []string{}
	dfs = func(left int) {
		if left == n {
			ret = min(ret, len(path))
			return
		}
		for i := left; i < n; i++ {
			if check(left, i, s) {
				path = append(path, s[left:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ret
}

func check(left, right int, s string) bool {
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func minCut(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}
	for i := 0; i < 2*n-1; i++ { // 中心拓展法
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			isPalindrome[l][r] = true
			l--
			r++
		}
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32 // 先设置为最大值,dp[0]初始值为0
	}
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			if isPalindrome[l][r] {
				dp[r+1] = min(dp[r+1], dp[l]+1)
			}
		}
	}
	return dp[n] - 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(minCut("ab"))
}
