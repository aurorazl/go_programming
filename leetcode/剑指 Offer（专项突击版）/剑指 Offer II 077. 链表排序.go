package main

/*
给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

归并排序
1.将链表分成两个子链表
2.对两个子链表排序后再将它们合并成一个排序的链表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(next)
	return mergeList(left, right)
}

func mergeList(a, b *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}
	return newHead.Next
}
