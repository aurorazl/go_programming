package main

import "fmt"

/*
由于互为字母异位词的两个字符串包含的字母相同，因此两个字符串中的相同字母出现的次数一定是相同的，
故可以将每个字母出现的次数使用字符串表示，作为哈希表的键。
*/
func groupAnagrams(strs []string) [][]string {
	storage := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := [26]int{}
		for _, char := range str {
			tmp[char-'a'] += 1
		}
		storage[tmp] = append(storage[tmp], str)
	}
	ans := [][]string{}
	for _, v := range storage {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
