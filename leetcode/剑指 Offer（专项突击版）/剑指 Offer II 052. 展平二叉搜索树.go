package main

/*
给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

先对输入的二叉搜索树执行中序遍历，将结果保存到一个列表中；
然后根据列表中的节点值，创建等价的只含有右节点的二叉搜索树，其过程等价于根据节点值创建一个链表。

在中序遍历的过程中改变节点指向。因为中序遍历，遍历到节点的顺序是刚好对应的。只要按照单链表的方式逐次连接即可。
*/

func increasingBST(root *TreeNode) *TreeNode {
	var inorder func(node *TreeNode)
	dummyHead := &TreeNode{}
	next := dummyHead
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		next.Right = node
		node.Left = nil
		next = node
		inorder(node.Right)
	}
	inorder(root)
	return dummyHead.Right
}
