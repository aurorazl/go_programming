package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"sort"
)

/*
给定一个可能有重复数字的整数数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。

回溯
*/

func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)
	var res [][]int
	var tmp []int
	var dfs func(idx, remained int)
	dfs = func(idx, remained int) {
		if remained == 0 {
			res = append(res, append([]int{}, tmp...))
			return
		}
		if idx == n {
			return
		}
		if candidates[idx] > remained { //要求必须排序，剪枝减少时间复杂度
			return
		}
		//jump_step := 1
		//for idx+jump_step<len(candidates) && candidates[idx+jump_step]==candidates[idx] {
		//	jump_step++
		//}
		//dfs(idx+jump_step, remained)					//跳过重复的开始位置，选择开始位置，重复的只要第一次那个
		//tmp = append(tmp, candidates[idx])
		//dfs(idx+1, remained-candidates[idx])	// 不重复求解，每个元素都尝试
		//tmp = tmp[:len(tmp)-1]

		for i := idx; i < n; i++ {
			if candidates[i] > remained {
				return
			}
			if i > idx && candidates[i-1] == candidates[i] {
				continue //忽略重复的开始位置
			}
			tmp = append(tmp, candidates[i])
			dfs(i+1, remained-candidates[i])
			tmp = tmp[:len(tmp)-1]
		}

	}
	dfs(0, target)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
