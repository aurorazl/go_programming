package main

import "fmt"

/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sum := 0
	var head *ListNode
	var cur *ListNode
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum = (x + y + carry)
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{sum, nil}
			cur = head
		} else {
			cur.Next = &ListNode{sum, nil}
			cur = cur.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return head
}
func main() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	l2 := &ListNode{9, &ListNode{9, nil}}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
