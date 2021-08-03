package main

import "fmt"

/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。
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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func recur(node *TreeNode, node2 *TreeNode) bool {
	if node2 == nil {
		return true
	}
	if node == nil || node.Val != node2.Val {
		return false
	}
	return recur(node.Left, node2.Left) && recur(node.Right, node2.Right)
}

func main() {
	a := loadTree([]interface{}{1, 2, 3, 4}, 0)
	b := loadTree([]interface{}{3}, 0)
	fmt.Println(isSubStructure(a, b))
}
