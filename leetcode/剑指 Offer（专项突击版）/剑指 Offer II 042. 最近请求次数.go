package main

/*

写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。
确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。

队列
*/

type RecentCounter struct {
	queue []int
	cnt   int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}, 0}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	this.cnt++
	for t-this.queue[0] > 3000 {
		this.queue = this.queue[1:]
		this.cnt--
	}
	return this.cnt
}
