package main

/*
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

双指针
哨兵节点主要是用来避免单独处理删除头节点的情况。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{-1, head}
	slow, fast := newHead, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
