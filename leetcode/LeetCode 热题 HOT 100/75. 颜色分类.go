package main

import "fmt"

/*
给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

1.单指针
	在第一次遍历中，我们将数组中所有的 0 交换到数组的头部。在第二次遍历中，我们将数组中所有的 1 交换到头部的 0 之后。
此时，所有的 22 都出现在数组的尾部，这样我们就完成了排序。

2.双指针
如果找到了 1，那么将其与 p1进行交换，并将 p1向后移动一个位置，这与方法一是相同的；
如果找到了 0，那么将其与 p0进行交换，并将 p0向后移动一个位置。再将i与p1位置替换（p1>p0）
*/

func sortColors(nums []int) {
	p0 := 0
	for index, i := range nums {
		if i == 0 {
			nums[index], nums[p0] = nums[p0], nums[index]
			p0++
		}
	}
	p1 := p0
	for i := p0; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p1 > p0 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p1++
			p0++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors2(a)
	fmt.Println(a)

}
