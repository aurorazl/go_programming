package LeetCode_热题_HOT_100

/*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]

nums 的长度恰好也为 nn，让 nums 充当哈希表
*/

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for _, i := range nums {
		target := (i - 1) % n //能保证对应位置+n
		nums[target] += n
	}
	for index, i := range nums {
		if i <= n {
			ans = append(ans, index+1)
		}
	}
	return ans
}
