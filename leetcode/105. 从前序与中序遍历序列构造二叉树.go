package main

/*
给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
对于任意一颗树而言，前序遍历的形式总是
	[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
	[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
而且数量相同
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
