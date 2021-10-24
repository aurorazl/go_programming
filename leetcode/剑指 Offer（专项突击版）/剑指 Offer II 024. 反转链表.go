package main

/*
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
