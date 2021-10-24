package main

/*
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。

深度优先搜索
在遇到 child 成员不为空的节点时，就要先去处理 child 指向的链表结构
需要将child 指向的链表结构进行扁平化，并且插入 node 与node 的下一个节点之间。

*/
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func dfs(node *Node) (last *Node) {
	cur := node
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			childLast := dfs(cur.Child)
			cur.Next = cur.Child
			cur.Child.Prev = cur
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}
			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}
