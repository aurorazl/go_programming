package main

import "fmt"

/*
给你一个整数数组 nums 和两个整数k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。

桶排序
按照元素的大小进行分桶，维护一个滑动窗口k内的元素对应的元素。
对于元素 x，其影响的区间为[x−t,x+t]。于是我们可以设定桶的大小为 t+1。如果两个元素同属一个桶，那么这两个元素必然符合条件。
如果两个元素属于相邻桶，那么我们需要校验这两个元素是否差值不超过 t。如果两个元素既不属于同一个桶，也不属于相邻桶，那么这两个元素必然不符合条件。
具体地，我们遍历该序列，假设当前遍历到元素 x，那么我们首先检查 x 所属于的桶是否已经存在元素，
如果存在，那么我们就找到了一对符合条件的元素，否则我们继续检查两个相邻的桶内是否存在符合条件的元素。

*/

func getId(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := map[int]int{}
	for i, num := range nums {
		id := getId(num, t+1)
		if _, ok := mp[id]; ok { //能分到同一个桶，而且滑动窗口内，肯定存在
			return true
		}
		if y, ok := mp[id-1]; ok && abs(num-y) <= t {
			return true
		}
		if y, ok := mp[id+1]; ok && abs(num-y) <= t {
			return true
		}
		mp[id] = num
		if i >= k {
			delete(mp, getId(nums[i-k], t+1))
		}
	}
	return false
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 1))
}
