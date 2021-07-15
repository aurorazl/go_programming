package main

import "fmt"

type treeNode struct {
	val interface{}
	left *treeNode
	right *treeNode
}
func FindTreeDepth(root *treeNode) int {
	if root == nil {
		return 0
	}
	first := []*treeNode{root}
	cnt := 0
	if len(first)==0 {
		return cnt
	}
	for len(first)>0 {
		var second []*treeNode
		for _,one := range first{
			if one.left!=nil{
				second = append(second, one.left)
			}
			if one.right != nil {
				second = append(second,one.right)
			}
		}
		cnt++
		first = second
	}
	return cnt
}

func main(){
	a := treeNode{val:1}
	b := treeNode{val:2}
	c := treeNode{val:2,left: &b,right: &a}
	fmt.Print(FindTreeDepth(&c))
}
