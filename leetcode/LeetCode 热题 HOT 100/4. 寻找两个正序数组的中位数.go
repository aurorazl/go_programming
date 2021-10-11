package main

import (
	"fmt"
	"math"
)

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	left := 0
	right := m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		fmt.Println(i, j, left, m, n)
		num_im := math.MinInt32
		if i != 0 {
			num_im = nums1[i-1]
		}
		num_i := math.MaxInt32
		if i != m {
			num_i = nums1[i]
		}
		num_jm := math.MinInt32
		if j != 0 {
			num_jm = nums2[j-1]
		}
		num_j := math.MaxInt32
		if j != n {
			num_j = nums2[j]
		}

		if num_im <= num_j {
			median1 = max(num_im, num_jm)
			median2 = min(num_i, num_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 {
		if n%2 == 0 {
			return float64(nums2[n/2]+nums2[n/2-1]) / 2.0
		} else {
			return float64(nums2[n/2])
		}

	}
	if m > n {
		return findMedianSortedArrays2(nums2, nums1)
	}
	left := 0
	right := m
	i := 0
	j := 0
	last_i, last_j := i, j
	for left <= right {
		i = (left + right) / 2
		j = (m+n+1)/2 - i
		fmt.Println(i, j, left, right, m, n)
		if j < n && i > 0 {
			if nums1[i-1] <= nums2[j] {
				last_i = i
				last_j = j
				left = i + 1
			} else {
				right = i - 1
			}
		} else {
			if nums1[i] <= nums2[j-1] {
				left = i + 1
			} else {
				last_i = i
				last_j = j
				right = i - 1
			}
		}

	}
	num_im := math.MinInt32
	num_i := math.MaxInt32
	num_jm := math.MinInt32
	num_j := math.MaxInt32
	i = last_i
	j = last_j
	if j != 0 {
		num_jm = nums2[j-1]
	}
	if i != 0 {
		num_im = nums1[i-1]
	}
	if i != m {
		num_i = nums1[i]
	}
	if j != n {
		num_j = nums2[j]
	}
	fmt.Println(num_im, num_jm, num_i, num_j)
	if (m+n)%2 == 0 {
		return float64(max(num_im, num_jm)+min(num_i, num_j)) / 2.0
	}
	return float64(max(num_im, num_jm))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func main() {
	fmt.Println(findMedianSortedArrays2([]int{4, 5}, []int{1, 2, 3}))
}
