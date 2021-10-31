package main

/*
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：

BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。

可以假设next()调用总是有效的，也就是说，当调用 next()时，BST 的中序遍历中至少存在一个下一个数字。

题意有点类似展平二叉搜索树

方法一：扁平化，递归获取中序遍历的全部结果并保存在数组中
方法二：迭代，实时维护栈。（而不是一次性获取出来）
*/

type BSTIterator struct {
	head *TreeNode
	cur  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var inorder func(node *TreeNode)
	dummyHead := &TreeNode{}
	next := dummyHead
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		next.Right = node
		node.Left = nil
		next = node
		inorder(node.Right)
	}
	inorder(root)
	return BSTIterator{dummyHead.Right, dummyHead.Right}
}

func (this *BSTIterator) Next() int {
	node := this.cur
	this.cur = this.cur.Right
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.cur == nil {
		return false
	}
	return true
}
