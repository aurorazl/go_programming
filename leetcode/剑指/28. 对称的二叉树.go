package main

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}
func dfs(node *TreeNode, node2 *TreeNode) bool {
	if node == nil && node2 == nil {
		return true
	}
	if node == nil || node2 == nil || node.Val != node2.Val {
		return false
	}
	return dfs(node.Right, node2.Left) && dfs(node.Left, node2.Right)
}
