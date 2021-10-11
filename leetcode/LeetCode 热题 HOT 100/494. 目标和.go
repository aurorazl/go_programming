package LeetCode_热题_HOT_100

/*
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

1.回溯
2.动态规划

*/

func findTargetSumWays(nums []int, target int) int {
	cnt := 0
	var backTrack func(int, int)
	backTrack = func(start int, sum int) {
		if start == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}
		backTrack(start+1, sum+nums[start])
		backTrack(start+1, sum-nums[start])
	}
	backTrack(0, 0)
	return cnt
}
