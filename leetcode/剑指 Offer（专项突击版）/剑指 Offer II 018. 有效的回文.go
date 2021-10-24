package main

import (
	"fmt"
	"unicode"
)

/*
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
本题中，将空字符串定义为有效的 回文串 。

双指针
*/

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) &&
			!unicode.IsDigit(rune(s[left])) {
			left++
		} else if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		} else {
			if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
				return false
			} else {
				left++
				right--
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
