package main

import "fmt"

/*
给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
*/
func maxProduct(words []string) int {
	n := len(words)
	ret := 0
	wordCode := make([]int, n)
	for i, word := range words {
		for _, char := range word {
			wordCode[i] |= 1 << (char - 'a') // 1 + 2 + 4 + 。。。。每个数字都是单独的，缺任何一个都可以看出来。对于二进制1的位置
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if wordCode[i]&wordCode[j] == 0 {
				ret = max(ret, len(words[i])*len(words[j]))
			}
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ret := 0
	for _, s := range "abc" {
		ret |= 1 << (s - 'a')
	}
	fmt.Println(ret)
	fmt.Println(maxProduct([]string{"aa", "b"}))
}
