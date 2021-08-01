package main

import "fmt"

/*
翻转一棵二叉树。
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	p := []*TreeNode{root}
	q := []*TreeNode{}

	for i := 0; len(p) != 0; i++ {
		res = append(res, []int{})
		for _, one := range p {
			res[i] = append(res[i], one.Val)
			if one.Left != nil {
				q = append(q, one.Left)
			}
			if one.Right != nil {
				q = append(q, one.Right)
			}
		}
		p = q
		q = []*TreeNode{}
	}
	return res
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}

func main() {
	a := loadTree([]interface{}{4, 2, 7, 1, 3, 6, 9}, 0)
	invertTree2(a)
	fmt.Println(levelOrder(a))
}
