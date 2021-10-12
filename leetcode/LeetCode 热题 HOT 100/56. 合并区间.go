package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})
	res := [][]int{}
	for _, i := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < i[0] {
			res = append(res, i)
		} else {
			res[len(res)-1][1] = max(i[1], res[len(res)-1][1])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(merge([][]int{{1, 2}, {2, 4}}))
}
