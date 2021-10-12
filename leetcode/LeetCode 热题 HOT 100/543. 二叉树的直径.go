package LeetCode_热题_HOT_100

import "go_programming/leetcode/LeetCode 热题 HOT 100"

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
*/

func diameterOfBinaryTree(root *LeetCode_热题_HOT_100.TreeNode) int {
	maxV := 0
	var dfs func(node *LeetCode_热题_HOT_100.TreeNode) int
	dfs = func(node *LeetCode_热题_HOT_100.TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		maxV = max(maxV, l+r+1)
		return max(l, r) + 1 //返回该节点为根的子树的深度
	}
	dfs(root)
	return maxV - 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
