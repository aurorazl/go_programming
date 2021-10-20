package main

/*
给定一棵二叉搜索树，请找出其中第k大的节点。

中序遍历倒序 为 递减序列 。(不然顺序还得遍历两次，才能找到n-k的位置)。
递归
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	ret := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if k == 0 {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			ret = node.Val
			return
		}
		dfs(node.Left)
	}
	dfs(root)
	return ret
}
