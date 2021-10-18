package main

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
方法一：递归
方法二：辅助栈（或队列）//每遍历一个节点，将子节点放入队列后，交换引用
	node.left, node.right = node.right, node.left
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
