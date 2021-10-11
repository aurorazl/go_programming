package LeetCode_热题_HOT_100

import "sort"

/*
假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
请你重新构造并返回输入数组people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

*/
func reconstructQueue(people [][]int) [][]int {
	st := make(map[int][]int)
	for _, one := range people {
		st[one[0]] = append(st[one[0]], one[1])
	}
	height := []int{}
	for k, _ := range st {
		height = append(height, k)
	}
	sort.Ints(height)
	for i := len(height) - 1; i >= 0; i-- {
		order := st[height[i]]
	}
}
