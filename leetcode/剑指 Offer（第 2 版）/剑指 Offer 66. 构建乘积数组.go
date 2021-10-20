package main

import "fmt"

/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。
不能使用除法。（不能先所有的相乘，再逐个除以）

两次遍历，第一次先算每个元素的左边的乘积（越来越多）
第二次再算每个元素的右边的乘积，两者相乘得到最终结果 （倒序，方便越来越多）

*/

func constructArr(a []int) []int {
	n := len(a)
	b := make([]int, n)
	if n == 0 {
		return b
	}
	b[0] = 1
	b[n-1] = 1
	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	tmp := 1
	for i := n - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
