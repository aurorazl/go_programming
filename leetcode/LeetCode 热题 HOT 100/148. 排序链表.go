package main

import "fmt"

/*
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
进阶：
	你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	newHead := &ListNode{0, head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := newHead, newHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return newHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	newHead := &ListNode{0, nil}
	cur := newHead
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			cur.Next = head2
			head2 = head2.Next
		} else {
			cur.Next = head1
			head1 = head1.Next
		}
		cur = cur.Next
	}
	for head1 != nil {
		cur.Next = head1
		cur = cur.Next
		head1 = head1.Next
	}
	for head2 != nil {
		cur.Next = head2
		cur = cur.Next
		head2 = head2.Next
	}
	return newHead.Next
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
	//a.Next.Next.Next.Next = a.Next.Next
	b := sortList(a)
	fmt.Println(b)
}
