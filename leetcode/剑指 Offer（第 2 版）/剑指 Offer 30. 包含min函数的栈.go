package main

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
建立辅助栈
*/

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 {
		this.min = append(this.min, x)
	} else {
		lastMin := this.min[len(this.min)-1]
		this.min = append(this.min, min(lastMin, x))
	}
	this.data = append(this.data, x)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}
