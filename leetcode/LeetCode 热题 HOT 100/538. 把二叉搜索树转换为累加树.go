package LeetCode_热题_HOT_100

import "go_programming/leetcode/LeetCode 热题 HOT 100"

/*
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node的新值等于原树中大于或等于node.val的值之和。
1.反序中序遍历

*/
func convertBST(root *LeetCode_热题_HOT_100.TreeNode) *LeetCode_热题_HOT_100.TreeNode {
	sum := 0
	var dfs func(*LeetCode_热题_HOT_100.TreeNode)
	dfs = func(node *LeetCode_热题_HOT_100.TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}
