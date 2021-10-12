package main

/*
给你两个单链表的头节点headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。


长短两个链表，后半部分相同，那么如果大家都走两轮，最终会在相交处相遇

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1 := headA
	node2 := headB
	for node1 != node2 {
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return node1
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
	a := loadListNode([]int{1, 2, 3})
	b := loadListNode([]int{4, 2, 3})
}
