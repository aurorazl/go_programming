package main

import "fmt"

/*
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。

*/

func isAlienSorted(words []string, order string) bool {
	array := [26]int{}
	for index, char := range order {
		array[char-'a'] = index
	}
	var isSort func(string, string) bool
	isSort = func(s, t string) bool {
		i := 0
		for ; i < len(s) && i < len(t); i++ {
			if array[s[i]-'a'] > array[t[i]-'a'] {
				return false
			}
			if array[s[i]-'a'] < array[t[i]-'a'] { // 某个单词小于，可以提前退出
				return true
			}
		}
		return i == len(s)
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		if !isSort(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}
