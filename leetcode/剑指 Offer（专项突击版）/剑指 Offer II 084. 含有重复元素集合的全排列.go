package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。

排序后，遍历的时候要看当前元素是否重复放置过了。
*/

func permuteUnique(nums []int) (res [][]int) {
	n := len(nums)
	sort.Ints(nums)
	vis := make([]bool, n)
	perm := []int{}
	var dfs func(int)
	dfs = func(cnt int) {
		if cnt == n {
			res = append(res, append([]int{}, perm...))
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] {
				continue
			}
			//if i>0 && nums[i-1] == nums[i] && vis[i-1]{     //确保重复的数字只能从后面倒序选，这样相等的数字有了可比性。 1 1 1 true false true这种情况也会走到
			//	continue
			//}
			if i > 0 && nums[i-1] == nums[i] && !vis[i-1] { //确保只能从前面顺序用，更节省时间，因为大多数情况下vis为false，剪枝更多， 1 1 1 true false true不会走到
				continue
			}
			vis[i] = true
			perm = append(perm, nums[i])
			dfs(cnt + 1)
			perm = perm[:len(perm)-1]
			vis[i] = false
		}
	}
	dfs(0)
	return
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 1}))
}
