package main

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value需要返回 -1

辅助的单调递减的双端队列：当一个元素进入队列的时候，它前面所有比它小的元素就不会再对答案产生影响。（最大值肯定不会从这些元素取）
4-3-2-1（每次移除元素的适合，判断队首元素是否是当前最大值，是就删除，不是就不删除）
递增的话适合栈
*/

type MaxQueue struct {
	q []int
	p []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	for len(this.p) > 0 && this.p[len(this.p)-1] <= value {
		this.p = this.p[:len(this.p)-1]
	}
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	ret := this.q[0]
	this.q = this.q[1:]
	if ret == this.p[0] {
		this.p = this.p[1:]
	}
	return ret
}
