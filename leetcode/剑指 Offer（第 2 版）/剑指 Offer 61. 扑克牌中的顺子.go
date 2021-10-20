package main

import (
	"fmt"
	"sort"
)

/*
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

方法一： 集合 Set + 遍历
	集合用来判断是否重复，除了0
	最大牌 - 最小牌 < 5 则可构成顺子

方法二：排序 + 遍历
	集合可以用排序来替代
	最大值、最小值也不用判断
*/

func isStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}
func main() {
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}
