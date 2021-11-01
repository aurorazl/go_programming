package main

/*
实现一个 MapSum 类，支持两个方法，insert和sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。

*/

type MapSum struct {
	children [26]*MapSum
	isNode   bool
	sum      int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	node := this
	for _, char := range key {
		trie := node.children[char-'a']
		if trie != nil {
			node = trie
		} else {
			node.children[char-'a'] = &MapSum{}
			node = node.children[char-'a']
		}
	}
	node.isNode = true
	node.sum = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this
	for _, char := range prefix {
		trie := node.children[char-'a']
		if trie == nil {
			return 0
		}
		node = trie
	}
	return getSum(node)
}

func getSum(node *MapSum) int {
	if node == nil {
		return 0
	}
	res := node.sum
	for _, child := range node.children {
		res += getSum(child)
	}
	return res
}
