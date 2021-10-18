package main

/*
给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

*/
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
