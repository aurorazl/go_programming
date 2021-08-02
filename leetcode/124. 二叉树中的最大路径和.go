package main

import (
	"fmt"
	"math"
)

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

空节点的最大贡献值等于 0。
非空节点的最大贡献值等于节点值与其子节点中的最大贡献值之和（对于叶节点而言，最大贡献值等于节点值）。
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
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftValue := max(maxGain(root.Left), 0)
		rightValue := max(maxGain(root.Right), 0)
		curValue := root.Val + leftValue + rightValue
		maxSum = max(maxSum, curValue)
		return root.Val + max(leftValue, rightValue)
	}
	maxGain(root)
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := loadTree([]interface{}{-10, 9, 20, nil, nil, 15, 7}, 0)
	fmt.Print(maxPathSum(a))
}
