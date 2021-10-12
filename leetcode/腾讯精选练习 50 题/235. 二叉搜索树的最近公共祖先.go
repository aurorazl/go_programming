package main

import "fmt"

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	ret := root
	for {
		if ret.Val > p.Val && ret.Val > q.Val {
			ret = ret.Left
		} else if ret.Val < p.Val && ret.Val < q.Val {
			ret = ret.Right
		} else {
			return ret
		}
	}
}

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

func main() {
	tree := loadTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 0)
	fmt.Println(lowestCommonAncestor(tree, tree.Left, tree.Left.Right))
}
