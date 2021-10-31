package main

import "fmt"

/*
给定两个以升序排列的整数数组 nums1 和 nums2,以及一个整数 k。
定义一对值(u,v)，其中第一个元素来自nums1，第二个元素来自 nums2。
请找到和最小的 k个数对(u1,v1), (u2,v2) ... (uk,vk)。

方法一："大根堆"(top k，缺点是堆顶最大元素i，j下标不能过滤其他下标)
由于 nums1 和 nums2 都是升序排序的，所以最小的k个数对肯定是在 nums1[0,k-1] 和 nums2[0,k-1] 中配对的，知道了这一点后，我们就可以开始解题了。
1.创建一个大根堆，堆中元素排序按照数对和进行逆序排序。
2.nums1 取前k个数（长度 n 小于 k 则取 n 个，nums2 相同）与数组2取前 k 个数组成 k*k 个数对。
3.每次都将当前配对的数对插入堆中，然后判断堆中元素总数是否超过 k , 如果超过 k 则将堆顶元素删除。
4.不断重复 3 操作，最后剩下的堆中的数对，就是和最小的 k 个数对。

方法二："小根堆"（不用遍历完K*K个数，堆顶下标为0，0，下一次肯定在i+1或者j+1里面）
大根堆的方法，其实还可以优化一下，在这里我们使用小根堆来解题，解题思路如下：

创建一个小根堆，小根堆中存放 <nums1对应元素的索引，nums2对应元素的索引>。
对小根堆进行初始化：将 <0,0>, ... , <k-1,0> 插入栈中（如果长度 n 小于 k 则取 n 个）。
选出和最小的数对【i,j】(堆顶)，将堆顶弹出，把 {nums1[i], nums2[j]} 保存到列表中。
同时将【i,j+1】插入堆中。
一直重复 3-4 过程，直至选到 k 个数对。

*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	sl := [][]int{}
	res := [][]int{}
	length := min(m, k)
	for i := 0; i < length; i++ {
		sl = append(sl, []int{i, 0})
	}
	var sift func(nums [][]int, low, high int)
	sift = func(nums [][]int, low, high int) {
		tmp := nums[low]
		i := low
		j := i*2 + 1
		for j <= high {
			if j < high && nums1[nums[j][0]]+nums2[nums[j][1]] > nums1[nums[j+1][0]]+nums2[nums[j+1][1]] { //最小堆
				j++
			}
			if nums1[nums[j][0]]+nums2[nums[j][1]] < nums1[tmp[0]]+nums2[tmp[1]] { //最小堆
				nums[i] = nums[j]
				i = j
				j = i*2 + 1
			} else {
				break
			}
		}
		nums[i] = tmp
	}
	for i := length/2 - 1; i >= 0; i-- {
		sift(sl, i, length-1)
	}
	for len(res) != k && len(sl) > 0 {
		minPair := sl[0]
		sl = sl[1:]
		res = append(res, []int{nums1[minPair[0]], nums2[minPair[1]]})
		if minPair[1]+1 < n {
			sl = append(sl, []int{})
			copy(sl[1:], sl[0:])
			sl[0] = []int{minPair[0], minPair[1] + 1}
			sift(sl, 0, len(sl)-1)
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
}
