package main

import "fmt"

/*
请为团队设计⼀个LRU公共组件，作为进程内缓存缓存使⽤。它需要⽀持的操作：
读： bool Get(key, result) - 如果关键字 (key) 存在于缓存中，则获取关键字的值，不存在则返
回相应false
写： void Put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插
⼊该组「关键字/值」。当缓存容量达到上限时，应该淘汰合适的Key以腾出空间。
*/
type TreeNode struct {
	key  interface{}
	val  interface{}
	prev *TreeNode
	next *TreeNode
}
type LRU struct {
	storage map[interface{}]*TreeNode
	size    int
	head    *TreeNode
	tail    *TreeNode
}

func NewLruInstance(size int) *LRU {
	return &LRU{
		storage: make(map[interface{}]*TreeNode, size),
		size:    size,
		head:    nil,
		tail:    nil,
	}
}

func (lru *LRU) Get(key interface{}) (val interface{}) {
	node, ok := lru.storage[key]
	if !ok {
		return
	}
	lru.MoveToHead(node)
	return node.val
}
func (lru *LRU) Set(key, val interface{}) {
	if node, ok := lru.storage[key]; ok {
		node.val = val
		lru.MoveToHead(node)
	}
	newNode := &TreeNode{key: key, val: val}
	lru.storage[key] = newNode
	lru.AddToHead(newNode)
	if len(lru.storage) > lru.size {
		tail := lru.RemoveTail()
		delete(lru.storage, tail.key)
	}
}
func (lru *LRU) RemoveTail() *TreeNode {
	tail := lru.tail
	lru.tail = tail.prev
	return tail
}
func (lru *LRU) MoveToHead(node *TreeNode) {
	if lru.tail == node {
		return
	}
	lru.RemoveNode(node)
	lru.AddToHead(node)
}
func (lru *LRU) AddToHead(node *TreeNode) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	} else {
		lru.head.prev = node
		node.next = lru.head
		lru.head = node
	}
}
func (lru *LRU) RemoveNode(node *TreeNode) {
	if lru.head == lru.tail {
		lru.head = nil
		lru.tail = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

func main() {
	lru := NewLruInstance(5)
	fmt.Println(lru.Get("111"))
	lru.Set("111", 222)
	lru.Set(111, 222)
	fmt.Println(lru.Get("111"))
	fmt.Println(lru.storage)
	lru.Set(222, 222)
	lru.Set(333, 222)
	lru.Set(444, 222)
	lru.Set(555, 222)
	lru.Set(666, 222)
	fmt.Println(lru.storage)
}
