package main

import "fmt"

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。
滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

单调队列
在队列中，这些下标按照从小到大的顺序被存储，并且它们在数组 nums 中对应的值是严格单调递减的。
由于队列中下标对应的元素是严格单调递减的，因此此时队首下标对应的元素就是滑动窗口中的最大值。
需要不断从队首弹出元素，直到队首元素在窗口中为止。
*/
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	ans := []int{}
	for i := 0; i < k; i++ {
		if len(queue) > 0 {
			for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
				queue = queue[:len(queue)-1]
			}
		}
		queue = append(queue, i)
	}
	ans = append(ans, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		ans = append(ans, nums[queue[0]])
	}
	return ans
}
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}
