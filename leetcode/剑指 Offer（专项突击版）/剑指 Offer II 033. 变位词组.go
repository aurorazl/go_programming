package main

import "sort"

/*
给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。

哈希表+排序key
*/

func groupAnagrams(strs []string) (res [][]string) {
	dict := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] > tmp[j]
		})
		newStr := string(tmp)
		dict[newStr] = append(dict[newStr], str)
	}
	for _, v := range dict {
		res = append(res, v)
	}
	return
}
