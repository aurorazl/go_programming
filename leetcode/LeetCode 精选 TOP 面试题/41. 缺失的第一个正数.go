package main

import "fmt"

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

思路：置换
实际上，对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1,N+1] 中。
这是因为如果 [1, N] 都出现了，那么答案是 N+1，否则答案是[1,N] 中没有出现的最小正整数。
以对数组进行一次遍历，对于遍历到的数 x=nums[i]，如果 x∈[1,N]，我们就知道 x 应当出现在数组中的 x−1 的位置，
因此交换 nums[i] 和nums[x−1]，这样 x 就出现在了正确的位置。在完成交换后，新的 nums[i] 可能还在 [1,N] 的范围内，我们需要继续进行交换操作，
直到 x∈[1,N]。

注意到上面的方法可能会陷入死循环。如果 nums[i] 恰好与nums[x−1] 相等，那么就会无限交换下去。
此时我们有 nums[i]=x=nums[x−1]，说明 x 已经出现在了正确的位置。因此我们可以跳出循环，开始遍历下一个数。

*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return n + 1
}
func main() {
	fmt.Print(firstMissingPositive([]int{1, 1}))
}
