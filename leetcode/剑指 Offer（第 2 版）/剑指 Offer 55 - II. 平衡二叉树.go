package main

import "math"

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

注意是每一个节点都要平衡

方法一：自顶向下的递归
	对于同一个节点，函数 height 会被重复调用，导致时间复杂度较高。(可以加缓存，以空间换时间)
	func height(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Left), height(root.Right)) + 1
	}
	return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
	或者return isBalanced(root.Left) && isBalanced(root.Right) && abs(height(root.Left),height(root.Right))<=1
	这两个顺序都会重复求值
方法二：自底向上的递归，先从子树开始判断。
	子树的高度，如果超出1的情况下，返回-1 否则正常返回。
	最后判断整棵树的高度是否大于等于0
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0 // 只要整棵树的高度是正常返回的即可。
}
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
