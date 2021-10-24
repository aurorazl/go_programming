package main

import "fmt"

/*
给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
注意：若s 和 t中每个字符出现的次数都相同且字符顺序不完全相同，则称s 和 t互为变位词（字母异位词）。
 仅包含小写字母

哈希
两个数组比较
一个数组 + 遍历diff(不能判断相同的字符串)

*/

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n || s == t {
		return false
	}
	array := [26]int{}
	for i := 0; i < n; i++ {
		array[s[i]-'a']++
		array[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if array[i] != 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isAnagram("nl", "cx"))
}
