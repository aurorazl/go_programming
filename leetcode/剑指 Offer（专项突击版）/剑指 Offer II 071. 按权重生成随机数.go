package main

import (
	"math/rand"
	"sort"
)

/*
给定一个正整数数组w ，其中w[i]代表下标 i的权重（下标从 0 开始），请写一个函数pickIndex，它可以随机地获取下标 i，选取下标 i的概率与w[i]成正比。
例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3)= 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3)= 0.75（即，75%）。
也就是说，选取下标 i 的概率为 w[i] / sum(w) 。

前缀和 + 二分查找

每个区间的左边界是在它之前出现的所有元素的和加上 1,右边界是到它为止的所有元素的和。
*/
type Solution struct {
	w []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.w[len(this.w)-1]) + 1
	return sort.SearchInts(this.w, x) //w是划分好的前缀和，二分查找
}
