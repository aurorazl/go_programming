package main

import "fmt"

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
注意条件的先后顺序
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val { //找到了其中一个
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { //最先返回的是最近包含pq的公共节点
		return root
	}
	if left != nil { //这里拿到的是底层传上来的最近祖先
		return left
	}
	return right
}

func main() {
	a := loadTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 0)
	fmt.Println(lowestCommonAncestor(a, a.Left, a.Right))
}
