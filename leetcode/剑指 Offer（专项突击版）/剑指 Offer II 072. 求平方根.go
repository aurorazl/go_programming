package main

/*
给定一个非负整数 x ，计算并返回 x 的平方根，即实现int sqrt(int x)函数。
正数的平方根有两个，只输出其中的正数平方根。
如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。

整数x的平方根一定小于或等于x
除0之外的所有整数的平方根都大于或等于1
整数x的平方根一定是在1到x的范围内，取这个范围内的中间数字mid，并判断mid的平方是否小于或等于x，
如果mid的平方小于x,那么接着判断(mid+1)的平方是否大于x，如果(mid+1)的平方大于x，那么mid就是x的平方根
如果mid的平方小于x并且(mid+1)的平方小于x，那么x的平方根比mid大，接下来搜索从mid+1到x的范围
如果mid的平方大于x，则x的平方根小于mid，接下来搜索1到mid-1的范围
然后在相应的范围内重复这个过程，总是取出位于范围中间的mid
*/

func mySqrt(x int) int {
	left := 1
	right := x
	for left <= right {
		mid := left + (right-left)/2
		if x/mid >= mid {
			if x/(mid+1) < mid+1 {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
