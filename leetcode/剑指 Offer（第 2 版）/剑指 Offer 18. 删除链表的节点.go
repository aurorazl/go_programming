package main

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
注意：此题对比原题有改动
*/
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head
}
