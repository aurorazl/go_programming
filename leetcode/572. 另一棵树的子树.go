package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root, sub *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, sub) || isSubtree(root.Left, sub) || isSubtree(root.Right, sub)
}

func check(root, sub *TreeNode) bool {
	if root == nil && sub == nil {
		return true
	}
	if root == nil || sub == nil {
		return false
	}
	if root.Val == sub.Val {
		return check(root.Left, sub.Left) && check(root.Right, sub.Right)
	}
	return false
}

func main() {
	a := &TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}
	b := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}
	fmt.Println(isSubtree(a, b))
}
