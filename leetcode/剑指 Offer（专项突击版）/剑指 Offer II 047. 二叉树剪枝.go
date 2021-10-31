package main

/*
给定一个二叉树 根节点root，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点（子树）的值为 0 的子树。
节点 node 的子树为node 本身，以及所有 node的后代。

从叶子节点逆推，不断往上删子树为0的节点

深度优先 + 后续遍历
*/

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Right == nil && root.Left == nil {
		root = nil
	}
	return root
}
