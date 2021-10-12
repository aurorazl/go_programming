package LeetCode_热题_HOT_100

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

遍历一次动态计算
*/
func maxProfit(prices []int) int {
	min_price := prices[0]
	maxValue := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
		} else {
			maxValue = max(maxValue, prices[i]-min_price)
		}
	}
	return maxValue
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
