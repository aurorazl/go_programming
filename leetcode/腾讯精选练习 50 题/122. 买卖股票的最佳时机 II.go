package main

/*
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

动态规划
定义状态 dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润，dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润（i 从 0 开始）。

*/
func maxProfit(prices []int) int {
	//dp := make([][2]int,len(prices))
	//dp[0][1] = -prices[0]
	dp0 := 0
	dp1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		//dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		//dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	//return dp[len(prices)-1][0]
	return dp0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
