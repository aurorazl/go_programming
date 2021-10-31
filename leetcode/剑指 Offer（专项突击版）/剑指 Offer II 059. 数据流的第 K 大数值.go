package main

/*
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
请实现 KthLargest类：
KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

优先队列(最小堆)

*/

type KthLargest struct {
	k  int
	sl []int
}

func Constructor(k int, nums []int) KthLargest {
	sl := []int{}
	for i := 0; i < len(nums); i++ {
		if i < k {
			sl = append(sl, nums[i])
			if i == k-1 {
				heapInit(sl)
			}
		} else {
			if nums[i] > sl[0] {
				sl[0] = nums[i]
				sift(sl, 0, k-1)
			}
		}
	}
	return KthLargest{k, sl}
}

func (this *KthLargest) Add(val int) int {
	if len(this.sl) < this.k {
		this.sl = append(this.sl, val)
		sift(this.sl, 0, len(this.sl)-1)
	} else {
		if val > this.sl[0] {
			this.sl[0] = val
			sift(this.sl, 0, this.k-1)
		}
	}
	return this.sl[0]
}

func sift(nums []int, low, high int) {
	tmp := nums[low]
	i := low
	j := i*2 + 1
	for j <= high {
		if j < high && nums[j] > nums[j+1] { //最小堆
			j++
		}
		if nums[j] < tmp { //最小堆
			nums[i] = nums[j]
			i = j
			j = i*2 + 1
		} else {
			break
		}
	}
	nums[i] = tmp
}

func heapInit(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		sift(nums, i, n-1)
	}
	//for j:=n-1;j>=0;j--{
	//	nums[0],nums[j] = nums[j],nums[0]
	//	sift(nums,0,j-1)
	//}
}
