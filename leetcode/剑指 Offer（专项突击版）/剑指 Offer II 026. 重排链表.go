package main

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0→ L1→ … → Ln-1→ Ln
请将其重新排列后变为：
L0→Ln→L1→Ln-1→L2→Ln-2→ …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

方法一：线性表
因为链表不支持下标访问，所以我们利用线性表存储该链表，然后利用线性表可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可。
方法二：寻找链表中点 + 链表逆序 + 合并链表(因为两链表长度相差不超过 11，因此直接合并即可。)

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
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func mergeList(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l1.Next = l2
		l2Next := l2.Next
		l1 = l1Next
		l2.Next = l1
		l2 = l2Next
	}
}
func reorderList(head *ListNode) {
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}
