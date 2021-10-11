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
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i > 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i-1] {
			j--
		}
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
