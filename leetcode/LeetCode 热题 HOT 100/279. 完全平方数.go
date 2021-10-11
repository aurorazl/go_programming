package main

import (
	"fmt"
	"math"
)

/*
给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

f[i]表示最少需要多少个数的平方来表示整数 i。
这些数必然落在区间 [1,n的开方]
我们可以枚举这些数，假设当前枚举到 j，那么我们还需要取若干数的平方，构成 i-j^2

*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minV := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minV = min(minV, dp[i-j*j]) //遍历过程中最小值可能在中间，比如4+4+4小于9+1+1+1
		}
		dp[i] = minV + 1
	}
	return dp[n]
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
func main() {
	fmt.Println(numSquares(12))
}
