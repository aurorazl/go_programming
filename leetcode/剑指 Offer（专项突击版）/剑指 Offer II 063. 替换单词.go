package main

import (
	"fmt"
	"strings"
)

/*
在英语中，有一个叫做词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为继承词(successor)。例如，词根an，跟随着单词other(其他)，可以形成新的单词another(另一个)。
现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
需要输出替换之后的句子。

思路：
初始化前缀树，将 dictionary 中的所有词根都插入前缀树中。
将句子 sentence 切分为一个个单词。
对每一个单词都进行查询：到前缀树中进行搜索，查看是否含有前缀树中的词根，如果有词根则将单词进行替换，如果没有词根则单词保持原样。
最后，把所有的单词拼接成句子返回（也可以查询完一个拼接一个）。

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

func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()
	for _, word := range dictionary {
		trie.Insert(word)
	}
	splits := strings.Split(sentence, " ")
	for i, split := range splits {
		for j := 0; j < len(split); j++ {
			if trie.Search(split[:j+1]) {
				splits[i] = split[:j+1]
				break
			}
		}
	}
	return strings.Join(splits, " ")
}

func main() {
	/*
		输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
		输出："the cat was rat by the bat"
	*/
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
