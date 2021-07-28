package main

import (
	"fmt"
	"math"
)

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
思路1：顺序合并
	前面2个先合并，在合并接下来的一个，直到为空
思路2：分治合并，将整个数组不断二分，直到最后只剩下两个，进行合并
思路3：优先队列合并
	维护当前每个链表没有被合并的元素的最前面一个，k 个链表就最多有 k 个满足这样条件的元素
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{-1, nil}
	cur := head
	nilCnt := 0
	for nilCnt < len(lists) {
		minValue := math.MaxInt32
		index := -1
		nilCnt = 0
		for i, one := range lists {
			if one == nil {
				nilCnt++
				continue
			}
			if one.Val < minValue {
				minValue = one.Val
				index = i
			}
		}
		if nilCnt < len(lists) {
			cur.Next = lists[index]
			cur = cur.Next
			lists[index] = lists[index].Next
		}
	}
	return head.Next
}

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	b := &ListNode{4, &ListNode{5, &ListNode{6, nil}}}
	c := &ListNode{7, &ListNode{8, &ListNode{9, nil}}}
	d := mergeKLists([]*ListNode{a, b, c})
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
