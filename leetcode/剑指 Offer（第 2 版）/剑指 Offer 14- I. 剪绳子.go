package main

import (
	"fmt"
)

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

递归，每个长度都有可能剪或者不剪。
动态规划
	每增加一米长度，之前的每一米处都有可能重新剪。
	两层循环
*/

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Println(i, j, dp)
			dp[i] = max(max(dp[i], j*(i-j)), j*dp[i-j])
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var cache = make(map[int]int, 0) // 缓存，不然超时
func cuttingRope2(n int) int {
	ans := 0
	if n == 2 || n == 1 {
		return 1
	}
	if cache[n-1] > 0 {
		return cache[n-1]
	}
	for i := 1; i < n-1; i++ {
		ans = max(max(ans, i*(n-i)), i*cuttingRope2(n-i)) // 要三种情况比较，不然会漏掉，比如3 = 1+2 这种情况
	}
	cache[n-1] = ans
	return ans
}

func main() {
	fmt.Println(cuttingRope2(31))
}
