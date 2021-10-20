package main

import "fmt"

/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

基础：
	每个骰子摇到 1 至 6 的概率相等
	将每个骰子的点数看作独立情况
	n 个骰子「点数和」的范围为 [n, 6n]，数量为6n−n+1=5n+1 种。

动态规划
概率 f(n - 1, x) 仅与 f(n,x+1) , f(n,x+2), ... , f(n,x+6) 相关。
*/

func dicesProbability(n int) []float64 {
	ret := make([]float64, 6)
	for i := 0; i < 6; i++ {
		ret[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ { // 第几枚骰子，前面已经初始化第一个
		tmp := make([]float64, 5*i+1)
		for j := 0; j < len(ret); j++ { // 对于之前的每个结果
			for k := 0; k < 6; k++ { // 当前骰子抛了第几个数
				tmp[j+k] += ret[j] / 6.0 // 每种可能的概率相加
			}
		}
		ret = tmp
	}
	return ret
}
func main() {
	fmt.Println(dicesProbability(2))
}
