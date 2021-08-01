package main

import "fmt"

/*
请判断一个链表是否为回文链表。
1->2->2->1

思路
双指针（复制链表值到数组列表中）或者找到中点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next

	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
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
	a := loadListNode([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindrome(a))
}
