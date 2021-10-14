package main

import "fmt"

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL。

思路：
	层次遍历
层次遍历基于广度优先搜索，它与广度优先搜索的不同之处在于，广度优先搜索每次只会取出一个节点来拓展，而层次遍历会每次将队列中的所有元素都拿出来拓展，
这样能保证每次从队列中拿出来遍历的元素都是属于同一层的，因此我们可以在遍历的过程中修改每个节点的 next 指针，同时拓展下一层的新队列。
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	stack := []*Node{root}
	var cur_stack []*Node
	for len(stack) > 0 {
		cur_stack = stack
		stack = nil
		for i, one := range cur_stack {
			if i+1 < len(cur_stack) {
				one.Next = cur_stack[i+1]
			}
			if one.Left != nil {
				stack = append(stack, one.Left)
			}
			if one.Right != nil {
				stack = append(stack, one.Right)
			}
		}
	}
	return root
}

func main() {
	var a []int
	a = append(a, 1)
	fmt.Println(a)
}
