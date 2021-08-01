package main

import "fmt"

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
  3
   / \
  9  20
    /  \
   15   7
[
  [3],
  [9,20],
  [15,7]
]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//双队列交替
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

func main() {
	fmt.Print(levelOrder(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
}
