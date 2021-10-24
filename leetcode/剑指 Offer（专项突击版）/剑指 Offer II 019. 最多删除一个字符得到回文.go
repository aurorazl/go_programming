package main

import (
	"fmt"
)

/*
给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。

双指针法	如果不相同，在删除一个字符之后再比较其它的字符就能够形成一个回文

这题因为不用去掉非字母的字符，直接用字符串比较
*/

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) ||
				isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	n := len(s)
	if n%2 == 0 {
		return s[:n/2] == reverse(s[n/2:])
	} else {
		return s[:n/2] == reverse(s[n/2+1:])
	}
}
func reverse(s string) string {
	bytes := []byte(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
	}
	return string(bytes)
}
func main() {
	fmt.Println(validPalindrome("tcaac"))
}
