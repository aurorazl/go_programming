package main

/*
给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。

哈希 + 迭代中序
双指针
一个指针指向有序列表小的一边，一个指针指向有序列表大的一边，然后一直搜寻，这种想法主要有两种实现方式：
	使用递归将二叉搜索树的结果保存到一个列表中，再对列表进行双指针查找。
	通过一个二叉搜索树迭代器
*/

func findTarget(root *TreeNode, k int) bool {
	cache := map[int]bool{}
	stk := []*TreeNode{}
	cur := root
	for cur != nil || len(stk) != 0 {
		if cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		} else {
			cur = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if cache[k-cur.Val] {
				return true
			}
			cache[cur.Val] = true
			cur = cur.Right
		}
	}
	return false
}

func findTarget2(root *TreeNode, k int) bool {
	cache := map[int]bool{}
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if cache[k-node.Val] {
			return true
		}
		cache[node.Val] = true
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
