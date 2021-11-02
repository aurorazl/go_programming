package main

/*
给定一个链表数组，每个链表都已经按升序排列。
请将所有链表合并到一个升序链表中，返回合并后的链表。

归并思路合并排序链表
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists)-1)
}
func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	a := mergeLists(lists, left, mid)
	b := mergeLists(lists, mid+1, right)
	return mergeList(a, b)
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
