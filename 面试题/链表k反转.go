package main

import "fmt"

/*
将给出的链表中的节点每\ k k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是\ k k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。
要求空间复杂度 \ O(1) O(1)
例如：
给定的链表是1\to2\to3\to4\to51→2→3→4→5
对于 \ k = 2 k=2, 你应该返回 2\to 1\to 4\to 3\to 52→1→4→3→5
对于 \ k = 3 k=3, 你应该返回 3\to2 \to1 \to 4\to 53→2→1→4→5
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	slow := head
	fast := head
	var newHead *ListNode
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	cur := slow.Next
	prev := slow
	last_prev := slow
	for fast.Next != nil {
		for i := 1; i < k; i++ {
			if fast.Next == nil {
				return newHead
			}
			fast = fast.Next
		}
		for i := 1; i < k; i++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		prev = last_prev
		last_prev = cur
		if newHead == nil {
			newHead = cur
		}
	}
	return newHead
}

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reverseKGroup(a, 2)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}
