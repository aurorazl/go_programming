package main

import "fmt"

type node3 struct {
	val  int
	next *node3
}

func singleChainSort(head *node3) *node3 {
	if head == nil || head.next == nil {
		return head
	}
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		fast, slow = fast.next.next, slow.next
	}
	mid := slow.next
	slow.next = nil
	left, right := singleChainSort(head), singleChainSort(mid)
	h := new(node3)
	res := h
	for left != nil && right != nil {
		if left.val < right.val {
			h.next = left
			left = left.next
		} else {
			h.next = right
			right = right.next
		}
		h = h.next
	}
	if left != nil {
		h.next = left
	} else {
		h.next = right
	}
	return res.next
}

func main() {
	a := &node3{4, &node3{2, &node3{3, nil}}}
	a = singleChainSort(a)
	for a != nil {
		fmt.Print(a.val)
		a = a.next
	}
}
