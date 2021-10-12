package main

import (
	"fmt"
	"strconv"
)

/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
解析：
	有点像全排列，不同的是每个数字对应多个字母，而且由先后顺序
	动态规划
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	storage := []string{""}
	phone := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for _, i := range digits {
		tmp := []string{}
		for _, one := range storage {
			num, _ := strconv.Atoi(string(i))
			for _, letter := range phone[num-2] {
				tmp = append(tmp, one+string(letter))
			}
		}
		storage = tmp
	}
	return storage
}

func main() {
	fmt.Println(letterCombinations("23"))
}
