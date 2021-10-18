package main

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorTree(root.Left, root.Right)
}

func isMirrorTree(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}
	return isMirrorTree(node1.Left, node2.Left) && isMirrorTree(node1.Right, node2.Right)
}
