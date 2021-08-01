package main

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

候选人(cand_num)初始化为nums[0]，票数count初始化为1。
当遇到与cand_num相同的数，则票数count = count + 1，否则票数count = count - 1。
当票数count为0时，更换候选人，并将票数count重置为1。
遍历完数组后，cand_num即为最终答案。
*/
func majorityElement(nums []int) int {
	cand_num := nums[0]
	cnt := 1
	for _, i := range nums[1:] {
		if i == cand_num {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				cand_num = i
				cnt = 1
			}
		}
	}
	return cand_num
}
