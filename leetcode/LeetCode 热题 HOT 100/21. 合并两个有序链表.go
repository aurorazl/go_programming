package main

import (
	"fmt"
)

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	b := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	c := mergeTwoLists(a, b)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}
