package main

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

先将所有区间按照起始位置排序，那么只需要比较相邻两个区间的结束位置就能知道它们是否重叠
如果它们重叠就将它们合并，然后判断合并的区间是否和下一个区间重叠
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ret := [][]int{}
	for _, one := range intervals {
		if len(ret) == 0 || ret[len(ret)-1][1] < one[0] {
			ret = append(ret, one)
		} else {
			ret[len(ret)-1][1] = max(one[1], ret[len(ret)-1][1])
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
