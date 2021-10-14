package main

import "math"

/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

方法一：寻找最大值
我们对数组 nums 进行一次遍历，找到最大值对应的位置即可。
方法二：迭代爬坡
在 [0, n) 的范围内随机一个初始位置 i，随后根据 nums[i−1],nums[i],nums[i+1] 三者的关系决定向哪个方向走：
如果 nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，我们可以直接返回 ii 作为答案；
如果 nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，我们需要往右走，即 i ←i+1；
如果 nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，我们需要往左走，即 i ←i−1；
如果 nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，我们可以朝任意方向走。
方法三：方法二的二分查找优化
对于当前可行的下标范围 [l,r]，我们随机一个下标 i；
如果下标 i 是峰值，我们返回 i 作为答案；
如果 nums[i]<nums[i+1]，那么我们抛弃 [l, i]的范围，在剩余 [i+1,r] 的范围内继续随机选取下标；
如果 nums[i]>nums[i+1]，那么我们抛弃 [i, r] 的范围，在剩余 [l,i−1] 的范围内继续随机选取下标。
*/

func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(a int) int {
		if a == -1 || a == n {
			return math.MinInt64
		}
		return nums[a]
	}
	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
