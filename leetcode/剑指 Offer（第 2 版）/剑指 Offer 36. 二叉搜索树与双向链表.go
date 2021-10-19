package main

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表(递增)。要求不能创建任何新的节点，只能调整树中节点指针的指向。

二叉搜索树的中序遍历为 递增序列。

中序遍历 + 返回递归当前子树的head tail
lHead, lTail := dfs(node.Left)
node
rHead, rTail := dfs(node.Right)
最后额外将root返回的head tail连成环即可。
*/
