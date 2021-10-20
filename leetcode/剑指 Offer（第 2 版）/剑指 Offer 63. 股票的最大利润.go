package main

import (
	"fmt"
	"math"
)

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

需要找出给定数组中两个数字之间的最大差值（即，最大利润）。此外，第二个数字（卖出价格）必须大于第一个数字（买入价格）。

一次遍历，找到min价格，作为买入价格，接下来不断尝试卖出
*/

func maxProfit(prices []int) int {
	buyPrice := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			if prices[i]-buyPrice > profit {
				profit = prices[i] - buyPrice
			}
		}
	}
	return profit
}
func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
