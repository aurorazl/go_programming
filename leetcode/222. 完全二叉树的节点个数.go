package leetcode

import "sort"

/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
完全二叉树 的定义如下：
在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。'
若最底层为第 h 层，则该层包含 1~2h个节点。

从根节点出发，每次访问左子节点，直到遇到叶子节点，该叶子节点即为完全二叉树的最左边的节点，经过的路径长度即为最大层数 hh。

    1            h = 0
   / \
  2   3          h = 1
 / \  /
4  5 6           h = 2，对于4（100），可以表示为0（000），5（101）可以表示为1（只向右一次001），6（110）可以表示为010

如何判断第 k 个节点是否存在呢？如果第 k个节点位于第 h 层，则 k 的二进制表示包含 h+1 位，其中最高位是 1，
其余各位从高到低表示从根节点到第 k 个节点的路径，0 表示移动到左子节点，1 表示移动到右子节点。
通过位运算得到第 k 个节点对应的路径，判断该路径对应的节点是否存在，即可判断第 k 个节点是否存在。

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false //忽略掉非叶子节点
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 { // 与运算，看k在当前层，应该走左还是右
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil // 返回能够使f(i)=true的最小的i，这个k也是所求的目标值
	}) - 1
}
