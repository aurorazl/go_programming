package main

import (
	"fmt"
	"math"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	storage := make(map[byte]int)
	length := len(s)
	maxCnt := 0
	left := 0
	for i := 0; i < length; i++ {
		if _, ok := storage[s[i]]; ok {
			left = int(math.Max(float64(storage[s[i]]+1), float64(left))) //确保重复的字符前的所有字符都不起作用了，双指针的思想
		}
		if i-left+1 > maxCnt {
			maxCnt = i - left + 1
		}
		storage[s[i]] = i
	}
	return maxCnt
}

func lengthOfLongestSubstring2(s string) int {
	storage := make(map[byte]int)
	length := len(s)
	maxCnt := 0
	left := 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(storage, s[i-1])
		}
		for left < length && storage[s[left]] == 0 { //确保删除了重复的字符
			storage[s[left]]++
			left++
		}
		if left-i > maxCnt {
			maxCnt = left - i
		}
	}
	return maxCnt
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("abba"))
}
