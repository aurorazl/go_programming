package main

import (
	"fmt"
)

/*
运用所掌握的数据结构，设计和实现一个 LRU (Least Recently Used，最近最少使用) 缓存机制 。

实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

哈希+双链表（单链表时间复杂度高）
*/

type Node struct {
	Val  int
	Prev *Node
	Next *Node
	Key  int
}
type LRUCache struct {
	dict map[int]*Node
	cap  int
	head *Node
	tail *Node
	size int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{make(map[int]*Node), capacity, head, tail, 0}
}

func (this *LRUCache) Get(key int) int {
	node, has := this.dict[key]
	if !has {
		return -1
	}
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, has := this.dict[key]
	if has {
		this.moveToHead(node)
		node.Val = value
	} else {
		newNode := &Node{Val: value, Key: key}
		this.AddToHead(newNode)
		this.dict[key] = newNode
		this.size++
		if this.cap < this.size {
			deleteNode := this.removeTail()
			delete(this.dict, deleteNode.Key)
			this.size--
		}
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	this.AddToHead(node)
}
func (this *LRUCache) AddToHead(node *Node) {
	this.head.Next.Prev = node
	node.Next = this.head.Next
	this.head.Next = node
	node.Prev = this.head
}
func (this *LRUCache) removeTail() *Node {
	node := this.tail.Prev
	node.Prev.Next = this.tail
	this.tail.Prev = node.Prev
	return node
}
func main() {
	lru := Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}
