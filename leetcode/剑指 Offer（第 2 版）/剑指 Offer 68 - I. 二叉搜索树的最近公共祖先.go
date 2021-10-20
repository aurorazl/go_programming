package main

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

与236 二叉树的最近公共祖先 不同，这里是二叉搜索树，条件更多了。
可以快速地找出树中的某个节点以及从根节点到该节点的路径

方法一：两次遍历
对于p，判断节点的值，判断大于小雨还是等于（找到了该节点）
对于节点 q 同理。在寻找节点的过程中，我们可以顺便记录经过的节点，这样就得到了从根节点到被寻找节点的路径。
当我们分别得到了从根节点到 p 和 q 的路径之后，我们就可以很方便地找到它们的最近公共祖先了。
一开始部分相同，遍历即可

方法二：一次遍历
我们从根节点开始遍历；
	如果当前节点的值大于 p 和 q 的值，说明 p 和 q 应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
	如果当前节点的值小于 p 和 q 的值，说明 p 和 q 应该在当前节点的右子树，因此将当前节点移动到它的右子节点；
	如果当前节点的值不满足上述两条要求，那么说明当前节点就是「分岔点」。
		此时，p 和 q 要么在当前节点的不同的子树中，要么其中一个就是当前节点。当前节点即时所求。

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	if ancestor == nil {
		return nil
	}
	for ancestor != nil {
		if ancestor.Val > p.Val && ancestor.Val > q.Val {
			ancestor = ancestor.Left
		} else if ancestor.Val < p.Val && ancestor.Val < q.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
	return
}
