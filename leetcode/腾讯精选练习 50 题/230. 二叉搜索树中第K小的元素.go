package main

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

通过构造 BST 的中序遍历序列，则第 k-1 个元素就是第 k 小的元素。

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func loadTree(nums []interface{}, index int) (root *TreeNode) {
	if index < len(nums) {
		if i, ok := nums[index].(int); ok {
			root = &TreeNode{i, nil, nil}
			root.Left = loadTree(nums, index*2+1)
			root.Right = loadTree(nums, index*2+2)
		}
	}
	return
}
func kthSmallest(root *TreeNode, k int) int {
	ret := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ret = append(ret, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ret[k-1]
}
