package main

import (
	"container/heap"
	"fmt"
)

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

方法一：堆
建立一个小顶堆，然后遍历「出现次数数组」：
如果堆的元素个数小于 k，就可以直接插入堆中。
如果堆的元素个数等于 k，则检查堆顶与当前出现次数的大小。如果堆顶更大，说明至少有 k 个数字的出现次数比当前值大，故舍弃当前值；否则，就弹出堆顶，并将当前值插入堆中。
遍历完成后，堆中的元素就代表了「出现次数数组」中前 k 大的值。
*/
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sift(nums [][2]int, low, high int) {
	tmp := nums[low]
	i := low
	j := i*2 + 1
	for j <= high {
		if j < high && nums[j][0] > nums[j+1][0] {
			j++
		}
		if nums[j][0] < tmp[0] {
			nums[i] = nums[j]
			i = j
			j = i*2 + 1
		} else {
			break
		}
	}
	nums[i] = tmp
}

func heapInit(nums [][2]int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		sift(nums, i, n-1)
	}
	//for j:=n-1;j>=0;j--{
	//	nums[0],nums[j] = nums[j],nums[0]
	//	sift(nums,0,j-1)
	//}
}
func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	seq := [][2]int{}
	for key, value := range occurrences {
		if len(seq) < k {
			seq = append(seq, [2]int{value, key})
			if len(seq) == k {
				heapInit(seq)
			}
		} else {
			if value > seq[0][0] {
				seq = append(append([][2]int{}, [2]int{value, key}), seq[1:]...)
				sift(seq, 0, len(seq)-1)
			}
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = seq[i][1]
	}
	return ret
}

func main() {
	fmt.Println(topKFrequent2([]int{5, 2, 5, 3, 5, 3, 1, 1, 3}, 2))
}
