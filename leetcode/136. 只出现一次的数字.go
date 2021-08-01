package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
字典
位运算
有以下三个性质。
	任何数和 00 做异或运算，结果仍然是原来的数，即 a \oplus 0=aa⊕0=a。
	任何数和其自身做异或运算，结果是 00，即 a \oplus a=0a⊕a=0。
	异或运算满足交换律和结合律，即 a \oplus b \oplus a=b \oplus a \oplus a=b \oplus (a \oplus a)=b \oplus0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
*/

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
