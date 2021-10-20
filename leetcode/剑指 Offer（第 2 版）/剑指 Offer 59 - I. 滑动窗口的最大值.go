package main

import "fmt"

/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

方法一：优先队列（最大堆）
每当我们向右移动窗口时，我们就可以把一个新的元素放入优先队列中，此时堆顶的元素就是堆中所有元素的最大值。
然而这个最大值可能并不在滑动窗口中，而是出现在滑动窗口左边界的左侧。
我们不断地移除堆顶的元素，直到其确实出现在滑动窗口中。

方法二：单调队列
维护一个严格单调递减的双向队列，存放索引下标
当滑动窗口向右移动时，
	新元素入队，与队尾元素不断比较，取代，直到满足单调队列
	不断从队首弹出元素，直到队首元素在窗口中为止。
*/

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	n := len(nums)
	if n == 0 {
		return q
	}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans := []int{}
	ans = append(ans, nums[q[0]])
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{}, 0))
}
