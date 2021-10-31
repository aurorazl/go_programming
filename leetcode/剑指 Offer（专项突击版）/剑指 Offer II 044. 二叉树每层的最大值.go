package main

import "math"

/*
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

层次遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		maxValue := math.MinInt32
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			maxValue = max(maxValue, node.Val)
		}
		res = append(res, maxValue)
		queue = queue[size:]
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
