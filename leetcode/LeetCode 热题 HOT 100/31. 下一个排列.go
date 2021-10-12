package main

import "fmt"

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须 原地 修改，只允许使用额外常数空间。
输入：nums = [1,2,3]
输出：[1,3,2]
思路：
	从右边开始，比较附近的两个数，看看是否升序，
	如果是升序，将峰值的左边较小值与峰值右边的较大值调换位置后，峰值位置后面的反转一下即可（之前是倒序的形式）
我们需要将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列。
同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	//找到开始升序的位置,
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i > 0 {
		j := n - 1
		// 找到右边的最小值作为较大数，比左边的较小数要大。
		for j >= 0 && nums[j] <= nums[i-1] {
			j--
		}
		//两者交换
		nums[i-1], nums[j] = nums[j], nums[i-1]
	}
	reverse(nums[i:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func main() {
	a := []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a)
}
