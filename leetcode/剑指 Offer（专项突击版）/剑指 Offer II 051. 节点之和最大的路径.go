package main

import "math"

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。

后续遍历+剪枝

*/
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))
		ans = max(ans, node.Val+left+right) //每颗子树可能为整颗树的路径最大和
		return node.Val + max(left, right)  //返回以root为顶点的树的最大"半径"和，这个不一定是最大值，但ans一定是
	}
	dfs(root)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
