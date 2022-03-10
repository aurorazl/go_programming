package main

import "fmt"

/*
翻转二叉树的步骤可总结如下：
（1）交换根结点的左子结点与右子结点；
（2）翻转根结点的左子树（递归调用当前函数）；
（3）翻转根结点的右子树（递归调用当前函数）。
非递归法：
遍历每个节点，将左右节点放入queue，再将左右节点互换。再从queue取出一个节点。
  	 4
   /   \
  2     7
 / \   / \
1   3 6   9

     4
   /   \
  7     2
 / \   / \
9   6 3   1

 */
type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func ReverseTree(root *TreeNode){
	if root == nil {
		return
	}
	root.left,root.right = root.right,root.left
	ReverseTree(root.left)
	ReverseTree(root.right)
}

func preOrderRecursion(root *TreeNode){
	if root == nil {
		return
	}
	fmt.Print(root.val,",")
	preOrderRecursion(root.left)
	preOrderRecursion(root.right)
}

func main(){
	a := TreeNode{2,&TreeNode{1,nil,nil},&TreeNode{3,nil,nil}}
	b := TreeNode{7,&TreeNode{6,nil,nil},&TreeNode{9,nil,nil}}
	c := TreeNode{4,&a,&b}
	preOrderRecursion(&c)
	ReverseTree(&c)
	fmt.Println()
	preOrderRecursion(&c)
}
