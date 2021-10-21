package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的正整数数组candidates和一个正整数target，找出candidates中所有可以使数字和为目标数target的唯一组合。
candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
对于给定的输入，保证和为target 的唯一组合数少于 150 个。
输入: candidates = [2,3,6,7], target = 7
输出: [[7],[2,2,3]]

动态规划(背包问题)
回溯
*/

func combinationSum(candidates []int, target int) [][]int {
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
		//if idx == n {
		//	return
		//}
		//if candidates[idx] > remained {		//要求必须排序，剪枝减少时间复杂度
		//	return
		//}
		//dfs(idx+1, remained)					// 这里为什么不减，是因为这里只是遍历开始位置，跳过当前元素（每一个位置可选可不选为开始位置）
		//tmp = append(tmp, candidates[idx])
		//dfs(idx, remained-candidates[idx])		// 这里不包含重复元素，否则会1 1 2重复求解。这里减的前后需要回溯。
		//tmp = tmp[:len(tmp)-1]

		for i := idx; i < n; i++ {
			if candidates[i] > remained {
				return
			}
			tmp = append(tmp, candidates[i])
			dfs(i, remained-candidates[i])
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, target)
	return res
}

//
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	dp := make([][][]int, target+1)
	for _, candidate := range candidates { // 避免去重问题，每一轮的这个数字之前没有使用过
		if candidate > target {
			break
		}
		dp[candidate] = append(dp[candidate], []int{candidate})
		for i := candidate; i <= target; i++ {
			for _, ans := range dp[i-candidate] {
				dp[i] = append(dp[i], append(append([]int{}, ans...), candidate))
			}
		}
	}
	return dp[target]
}

func combinationSum3(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	dp := make([][]int, target+1)
	for i := candidates[0]; i <= target; i++ {
		for _, candidate := range candidates {
			// 注意去重问题，(1+1+1)+2 (1+2)+1+1
			fmt.Println(candidate)
		}
	}
	return dp
}

func main() {
	fmt.Println(combinationSum([]int{1, 2}, 3))
}
