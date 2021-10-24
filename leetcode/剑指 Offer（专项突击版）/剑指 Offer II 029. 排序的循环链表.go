package main

/*
给定循环升序列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。


1.插入节点非最大值也非最小值	cur.val<=insertVal&&cur.next.val>=insertVal
2.当插入值是最小值或最大值时
	(1)插入值是最大值cur.val<=insertVal
	(2)插入值是最小值insertVal<=cur.next.val

*/
type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{x, nil}
		node.Next = node
		return node
	}
	cur := aNode
	for cur.Next != aNode { // 遍历一次即可
		if cur.Val <= x && cur.Next.Val >= x || (cur.Next.Val < cur.Val && (cur.Val <= x || cur.Next.Val >= x)) {
			break
		}
		cur = cur.Next
	}
	cur.Next = &Node{x, cur.Next}
	return aNode
}
