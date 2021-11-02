package main

/*
给定两个数组，arr1 和arr2，
arr2中的元素各不相同
arr2 中的每个元素都出现在arr1中
对 arr1中的元素进行排序，使 arr1 中项的相对顺序和arr2中的相对顺序相同。未在arr2中出现过的元素需要按照升序放在arr1的末尾。

自定义排序 + 哈希表存储arr2先后
计数排序
具体地，我们使用一个长度为 1001（下标从 0 到 1000）的数组 frequency，记录每一个元素在数组 arr1中出现的次数。
随后我们遍历数组 arr2，当遍历到元素 x 时，我们将 frequency[x] 个 x 加入答案中，并将 frequency[x] 清零。
当遍历结束后，所有在 arr2中出现过的元素就已经有序了。
此时还剩下没有在 arr2中出现过的元素，因此我们还需要对整个数组 frequency 进行一次遍历。
当遍历到元素 x 时，如果 frequency[x] 不为 0，我们就将 frequency[x] 个 x 加入答案中。

*/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}
