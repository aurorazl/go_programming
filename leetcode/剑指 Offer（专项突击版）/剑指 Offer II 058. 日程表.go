package main

import "fmt"

/*
请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，
注意，这里的时间是半开区间，即 [start, end), 实数x 的范围为， start <= x < end。
当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false并且不要将该日程安排添加到日历中。
请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)

队列存放[start, end)，book时可以适当合并，二分查找要book的下标。
*/

type MyCalendar struct {
	events [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	index := this.searchIndex(start)
	if index == -1 {
		return false
	}
	if (index > 0 && this.events[index-1][1] > start) || (index < len(this.events) && this.events[index][0] < end) {
		return false
	}
	this.events = append(this.events, [2]int{})
	copy(this.events[index+1:], this.events[index:])
	this.events[index] = [2]int{start, end}
	return true
}
func (this *MyCalendar) searchIndex(start int) int {
	left := 0
	right := len(this.events) - 1
	for left <= right {
		mid := (left + right) / 2
		if this.events[mid][0] > start {
			right = mid - 1
		} else if this.events[mid][0] < start {
			left = mid + 1
		} else {
			return -1
		}
	}
	return left
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(37, 50))
	fmt.Println(calendar.Book(33, 50))
	fmt.Println(calendar.Book(20, 33))
}
