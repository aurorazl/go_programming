package main

import (
	"fmt"
	"strconv"
)

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

动态规划，遍历数字字符串
字符串的第 i 位置：
	可以单独作为一位来翻译
	如果第 i - 1 位和第 i 位组成的数字在 10 到 25 之间，可以把这两位连起来翻译   ---》 动态方程

	由于第i位的结果，可以由i + i-1 和 i得来，类似台阶问题，可以用滚动数组优化。
*/

func translateNum(num int) int {
	s := strconv.Itoa(num)
	p, q, r := 0, 0, 1 // 初始状态
	for i := 0; i < len(s); i++ {
		p, q, r = q, r, 0 // r一定包含q的情况，但p的情况待定
		r += q
		if i == 0 {
			continue
		}
		pre := s[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

func main() {
	fmt.Println(translateNum(12258))
}
