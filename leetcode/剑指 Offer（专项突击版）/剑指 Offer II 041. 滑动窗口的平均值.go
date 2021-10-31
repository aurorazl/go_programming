package main

/*
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
实现 MovingAverage 类：
	MovingAverage(int size) 用窗口大小 size 初始化对象。
	double next(int val)成员函数 next每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，
		即滑动窗口里所有数字的平均值。
解法：队列
*/
type MovingAverage struct {
	queue []int
	cap   int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, size, 0}
}

func (this *MovingAverage) Next(val int) float64 {
	this.queue = append(this.queue, val)
	this.sum += val
	if len(this.queue) > this.cap {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	return float64(this.sum) / float64(len(this.queue))
}
