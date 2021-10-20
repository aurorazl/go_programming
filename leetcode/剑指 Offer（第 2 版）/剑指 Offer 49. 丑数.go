package main

import "fmt"

/*
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

方法二：动态规划

动态方程如何写？
定义三个指针 p_2,p_3,p_5 ，表示下一个丑数是当前指针指向的丑数乘以对应的质因数。这三个指针顺序不定。
后面的丑数一定由前面的丑数乘以2，或者乘以3，或者乘以5得来。
*/

func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, min(x3, x5))
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(11))
}
