package main

import "fmt"

/*
给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

深度优先搜索
穷举所有的可能，我们访问每一个节点 node，检测以 node 为起始节点且向下延深的路径有多少种。

前缀和
优化重复计算
假设当前从根节点 root 到节点 node 的前缀和为 curr，
则此时我们在已保存的前缀和查找是否存在前缀和刚好等于 curr−targetSum。假设从根节点 root 到节点 node 的路径中存在节点 pi
到根节点 root 的前缀和为 curr−targetSum，则节点 }pi+1到 node 的路径上所有节点的和一定为 targetSum。

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
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	var dfs func(node *TreeNode, left int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		val := node.Val
		if left == val { // 每个节点只会计算一次
			res++
		}
		dfs(node.Left, left-val)
		dfs(node.Right, left-val)
		return
	}
	if root == nil {
		return res
	}
	dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func pathSum2(root *TreeNode, targetSum int) int {
	ans := 0
	sum := map[int]int{0: 1}
	var dfs func(node *TreeNode, curr int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		ans += sum[curr-targetSum]
		sum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		sum[curr]-- //减掉
		return
	}
	dfs(root, 0)
	return ans
}
func main() {
	fmt.Println(pathSum2(loadTree([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}, 0), 8))
}
