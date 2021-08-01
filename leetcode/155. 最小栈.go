package main

import "math"

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
	push(x) —— 将元素 x 推入栈中。
	pop()—— 删除栈顶的元素。
	top()—— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。

只需要设计一个数据结构，使得每个元素 a 与其相应的最小值 m 时刻保持一一对应。因此我们可以使用一个辅助栈，与元素栈同步插入与删除，用于存储与每个元素对应的最小值。
当一个元素要入栈时，我们取当前辅助栈的栈顶存储的最小值，与当前元素比较得出最小值，将这个最小值插入辅助栈中；
当一个元素要出栈时，我们把辅助栈的栈顶元素也一并弹出；
在任意一个时刻，栈内元素的最小值就存储在辅助栈的栈顶元素中。

*/
type MinStack struct {
	stack     []int
	min_stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min_stack: []int{math.MaxInt64}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.min_stack = append(this.min_stack, min(this.min_stack[len(this.min_stack)-1], val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min_stack = this.min_stack[:len(this.min_stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min_stack[len(this.min_stack)-1]
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
