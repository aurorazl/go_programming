package main

import "fmt"

/*
给定一个无重复元素的正整数数组candidates和一个正整数target，找出candidates中所有可以使数字和为目标数target的唯一组合。
candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
对于给定的输入，保证和为target 的唯一组合数少于 150 个。
输入: candidates = [2,3,6,7], target = 7
输出: [[7],[2,2,3]]

动态规划
*/
func combinationSum(candidates []int, target int) [][]int {
	dp := [][][]int{}
	for i := 0; i <= target; i++ {
		dp = append(dp, [][]int{})
	}
	for _, i := range candidates {
		if i <= target {
			dp[i] = append(dp[i], []int{i})
		}
		for j := i + 1; j <= target; j++ {
			for _, one := range dp[j-i] {
				newOne := make([]int, len(one))
				copy(newOne, one)
				newOne = append(newOne, i)
				dp[j] = append(dp[j], newOne)
			}
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 10))
}
