package main

import "fmt"

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
快慢指针
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	slow, fast := newHead, head //避免了head本身要被移除的情况判断
	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	a = removeNthFromEnd(a, 5)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}
