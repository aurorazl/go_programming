package main

import "fmt"

/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

208.题目
*/
type Trie struct {
	children [26]*Trie
	isNode   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, char := range word {
		trie := node.children[char-'a']
		if trie != nil {
			node = trie
		} else {
			node.children[char-'a'] = &Trie{}
			node = node.children[char-'a']
		}
	}
	node.isNode = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isNode
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, char := range prefix {
		trie := node.children[char-'a']
		if trie == nil {
			return nil
		}
		node = trie
	}
	return node
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
}
