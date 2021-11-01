package main

import "sort"

/*
单词数组words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：

words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给定一个单词数组words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。

words = ["time", "me", "bell"]
一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。

后缀树
	当前单词是否属于某个单词的后缀 —> 单词倒序是否属于某个单词的前缀

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
func (this *Trie) Insert(word string) bool {
	node := this
	flag := false
	n := len(word)
	for i, _ := range word {
		trie := node.children[word[n-i-1]-'a']
		if trie != nil {
			node = trie
		} else {
			node.children[word[n-i-1]-'a'] = &Trie{}
			node = node.children[word[n-i-1]-'a']
			flag = true
		}
	}
	node.isNode = true
	return flag
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
func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	trie := Constructor()
	res := 0
	for _, word := range words {
		if trie.Insert(word) {
			res += len(word) + 1
		}
	}
	return res
}
