package main

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
层序遍历 + 队列 + (添加首部or尾部/不同方向打印)
*/
type TreeNode struct {
	Val   intLeft
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

}
