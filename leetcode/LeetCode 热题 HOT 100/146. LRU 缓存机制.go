package main

import "fmt"

/*
运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：
	LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
	int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
		当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
进阶：你是否可以在O(1) 时间复杂度内完成这两种操作
进阶：加上哨兵作为head、tail。
*/
type DLinkedNode struct {
	key, Val   int
	Prev, Next *DLinkedNode
}
type LRUCache struct {
	capacity int
	storage  map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity, make(map[int]*DLinkedNode, capacity),
		&DLinkedNode{0, 0, nil, nil}, &DLinkedNode{0, 0, nil, nil}}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}
func (this *LRUCache) Get(key int) int {
	node, ok := this.storage[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key, val int) {
	if node, ok := this.storage[key]; ok {
		node.Val = val
		this.moveToHead(node)
	} else {
		if len(this.storage) >= this.capacity {
			this.removeTail()
		}
		newNode := &DLinkedNode{key, val, nil, nil}
		this.storage[key] = newNode
		this.addHead(newNode)
	}
}

func (this *LRUCache) removeTail() {
	node := this.tail.Prev
	delete(this.storage, node.key)
	this.removeNode(node)

}
func (this *LRUCache) addHead(node *DLinkedNode) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addHead(node)
}
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
