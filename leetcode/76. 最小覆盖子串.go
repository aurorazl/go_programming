package main

import (
	"fmt"
	"math"
)

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

*/

func minWindow(s string, t string) string {
	storage := make(map[string]int)
	needCnt := len(t)
	for i, _ := range t {
		storage[string(t[i])]++
	}
	left := 0
	res := []int{0, math.MaxInt32}
	for index, _ := range s {
		if storage[string(s[index])] > 0 {
			needCnt--
		}
		storage[string(s[index])]--
		if needCnt == 0 { //当前满足
			for storage[string(s[left])] != 0 {
				storage[string(s[left])]++
				left++
			}
			if index-left < res[1]-res[0] {
				res = []int{left, index}
			}
			storage[string(s[left])]++
			needCnt++
			left++
		}
	}
	if res[1] > len(s) {
		fmt.Println(res)
		return ""
	}
	return s[res[0] : res[1]+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
