package main

/*
给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。

每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

深度优先搜索
广度优先搜索（队列，每次取出判断是否叶子节点，队列存放的是（node，num））
*/

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, num int) int
	dfs = func(node *TreeNode, num int) int {
		if node == nil {
			return 0
		}
		sum := num*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}
