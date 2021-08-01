package main

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
动态规划
*/
func permute(nums []int) [][]int {
	storage := [][]int{}
	for _, i := range nums {
		if len(storage) == 0 {
			storage = append(storage, []int{i})
		} else {
			tmp := [][]int{}
			for _, j := range storage {
				j_copy := make([]int, len(j))
				for index := 0; index <= len(j); index++ {
					copy(j_copy, j)
					right := append([]int{i}, j_copy[index:]...)
					new_one := append(j_copy[0:index], right...)
					tmp = append(tmp, new_one)
				}
			}
			storage = tmp
		}
	}
	return storage
}
