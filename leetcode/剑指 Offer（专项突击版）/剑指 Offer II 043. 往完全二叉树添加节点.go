package main

/*
完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1个节点）的，并且所有的节点都尽可能地集中在左侧。
设计一个用完全二叉树初始化的数据结构CBTInserter，它支持以下几种操作：
	CBTInserter(TreeNode root)使用根节点为root的给定树初始化该数据结构；
	CBTInserter.insert(int v) 向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
	CBTInserter.get_root() 将返回树的根节点。

注意题目要求返回新节点的父节点，需要找到第一个缺左或右节点的节点（层次遍历）

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	for len(queue) != 0 && queue[0].Right != nil && queue[0] != nil {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return CBTInserter{root, queue}
}

func (this *CBTInserter) Insert(v int) int {
	node := this.queue[0] // 队首维护缺子节点的
	if node.Left == nil {
		node.Left = &TreeNode{Val: v}
	} else {
		node.Right = &TreeNode{Val: v}
		this.queue = append(this.queue, node.Left)
		this.queue = append(this.queue, node.Right)
		this.queue = this.queue[1:]
	}
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
