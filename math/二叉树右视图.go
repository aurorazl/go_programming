package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func foreachNodeList(x []*TreeNode) (nextNode []*TreeNode) {
	for _, one := range x {
		if one.Left != nil {
			nextNode = append(nextNode, one.Left)
		}
		if one.Right != nil {
			nextNode = append(nextNode, one.Right)
		}
	}
	return
}
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	nextNode := []*TreeNode{root}
	for root != nil {
		ret = append(ret, root.Val)
		nextNode = foreachNodeList(nextNode)
		if len(nextNode) == 0 {
			break
		}
		root = nextNode[len(nextNode)-1]
	}
	return ret
}
func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Printf("Hello World!\n")
	ret := make([]TreeNode, 0)
	ret = []TreeNode{nil}
	//ret = append(ret,nil)
	fmt.Print(ret)
}
