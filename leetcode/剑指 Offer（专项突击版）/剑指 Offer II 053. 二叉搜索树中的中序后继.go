package main

import "fmt"

/*
给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
节点p的后继是值比p.val大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。

深度优先搜索方法是可以找到节点p的位置的，非递归的中序遍历有栈，找到节点p后设置flag，下一步就是中序后继。

方法2：
	二叉搜索树是有序的
	如果当前节点的值小于或等于节点p的值，那么节点p的下一个节点应该在它的右子树
	如果当前节点的值大于节点p的值，那么当前节点有可能是p的下一个节点，先存着，再去遍历左子树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	stk := []*TreeNode{}
	cur := root
	flag := false
	for cur != nil || len(stk) != 0 {
		if cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		} else {
			cur = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if flag {
				break
			}
			if cur.Val == p.Val {
				flag = true
			}
			cur = cur.Right
		}
	}
	return cur
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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	cur := root
	var res *TreeNode
	for cur != nil {
		if cur.Val > p.Val {
			res = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return res
}
func main() {
	fmt.Println(inorderSuccessor(loadTree([]interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 0), &TreeNode{Val: 6}))
}
