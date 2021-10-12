package main

import "fmt"

/*
给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。
你是否可以使用 O(1) 空间解决此题？

第一次相遇时，假设慢指针 slow 走了 k 步，那么快指针 fast 一定走了 2k 步，也就是说比 slow 多走了 k 步（也就是环的长度）。
设相遇点距环的起点的距离为 m，那么环的起点距头结点 head 的距离为 k - m，也就是说如果从 head 前进 k - m 步就能到达环起点。
所以，只要我们把快慢指针中的任一个重新指向 head，然后两个指针同速前进，k - m 步后就会相遇，相遇之处就是环的起点了。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head //和判断是否有环不同，大家同一起始位置，不然公式不成立
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
func loadListNode(nums []int) (head *ListNode) {
	var cur *ListNode
	for _, i := range nums {
		if head == nil {
			head = &ListNode{i, nil}
			cur = head
		} else {
			cur.Next = &ListNode{i, nil}
			cur = cur.Next
		}
	}
	return
}
func main() {
	a := loadListNode([]int{3, 2, 0, -4})
	a.Next.Next.Next.Next = a.Next
	fmt.Println(detectCycle(a))
}
