package main

import (
	"fmt"
)

/*
给定一个二叉树，检查它是否是镜像对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

//迭代
/*
它们的两个根结点具有相同的值
每个树的右子树都与另一个树的左子树镜像对称
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricReverse(root *TreeNode) bool {
	return check(root.Left, root.Right)
}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Right, q.Left) && check(p.Left, q.Right)
}

//迭代
func isSymmetric(root *TreeNode) bool {
	stack := []*TreeNode{}
	p, q := root, root
	stack = append(stack, p)
	stack = append(stack, q)
	for len(stack) != 0 {
		p, q = stack[0], stack[1]
		stack = stack[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if q.Val != p.Val {
			return false
		}
		stack = append(stack, q.Right)
		stack = append(stack, p.Left)
		stack = append(stack, q.Left)
		stack = append(stack, p.Right)
	}
	return true
}
func main() {
	fmt.Print(isSymmetric(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}))
}
