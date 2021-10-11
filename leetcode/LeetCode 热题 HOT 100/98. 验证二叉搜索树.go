package main

import (
	"fmt"
	"math"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
*/

//递归
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	lastMin := math.MinInt64
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if lastMin >= root.Val {
			return false
		}
		lastMin = root.Val
		root = root.Right
	}
	return true
}

func main() {
	fmt.Println(isValidBST(&TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}}}))
}
