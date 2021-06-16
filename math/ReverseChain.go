package main

type node struct {
	value    int
	nextNode *node
}

func reverseNode(head *node) *node {
	//  先声明两个变量
	//  前一个节点
	var preNode, nextNode *node
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.nextNode
		//  将头节点指向前一个节点
		head.nextNode = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}
