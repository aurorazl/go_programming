package main

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

后序遍历定义： [ 左子树 | 右子树 | 根节点 ] ，即遍历顺序为 “左、右、根” 。
中序遍历：[ 左子树 | 根节点 ｜ 右子树 ]
前序遍历：[根节点 | 左子树 ｜ 右子树]

1。先根据这个前提找到左右子树
2。由于左区间遍历的时候都满足《根，接下来只要判断右区间都满足》根即可
3。递归再去判断左右子树是否满足
*/

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n == 0 {
		return true
	}
	rootVal := postorder[n-1]
	i := 0
	for ; postorder[i] < rootVal; i++ {
	}
	for j := i; j < n-1; j++ {
		if postorder[j] < rootVal {
			return false
		}
	}
	return verifyPostorder(postorder[:i]) && verifyPostorder(postorder[i:n-1])
}
