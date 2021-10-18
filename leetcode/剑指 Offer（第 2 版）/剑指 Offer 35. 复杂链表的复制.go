package main

import "fmt"

/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

回溯 + 哈希表缓存(不然会死循环)
	当我们拷贝节点时，「当前节点的随机指针指向的节点」可能还没创建
	利用回溯的方式，让每个节点的拷贝操作相互独立。
	如果这两个节点中的任何一个节点的新节点没有被创建，我们都立刻递归地进行创建。当我们拷贝完成，回溯到当前层时，我们即可完成当前节点的指针赋值。
*/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cachedNode := map[*Node]*Node{}
	var deepCopy func(head *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, has := cachedNode[node]; has {
			return n
		}
		next := &Node{}
		cachedNode[node] = next
		next.Val = node.Val
		next.Next = deepCopy(node.Next)
		next.Random = deepCopy(node.Random)
		return next
	}
	return deepCopy(head)
}

func main() {
	a := &Node{7, nil, nil}
	b := &Node{13, nil, nil}
	c := &Node{11, nil, nil}
	d := &Node{10, nil, nil}
	e := &Node{1, nil, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	b.Random = a
	c.Random = e
	d.Random = c
	e.Random = a
	fmt.Println(copyRandomList(a))
}
