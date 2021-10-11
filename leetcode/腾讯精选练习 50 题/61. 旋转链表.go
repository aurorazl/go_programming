package main

import "fmt"

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
思路：
记给定链表的长度为 n，注意到当向右移动的次数 k≥n 时，我们仅需要向右移动 k mod n 次即可。因为每 n 次移动都会让链表变为原状。
这样我们可以知道，新链表的最后一个节点为原链表的第(n−1)−(k mod n) 个节点（从 0 开始计数）。
这样，我们可以先将给定的链表连接成环，然后将指定位置断开。
具体代码中，我们首先计算出链表的长度 nn，并找到该链表的末尾节点，将其与头节点相连。这样就得到了闭合为环的链表。然后我们找到新链表的最后一个节点（即原链表的第 (n - 1) - (k \bmod n)(n−1)−(kmodn) 个节点），将当前闭合为环的链表断开，即可得到我们所需要的结果。
特别地，当链表长度不大于 11，或者 kk 为 nn 的倍数时，新链表将与原链表相同，我们无需进行任何处理。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	iter.Next = head
	add := n - k%n - 1
	iter = head // 两步可以缩减为add := n-k%n
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
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
	fmt.Println(rotateRight(loadListNode([]int{1, 2, 3, 4, 5}), 2))
}
