package LeetCode_热题_HOT_100

/*
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
	Trie() 初始化前缀树对象。
	void insert(String word) 向前缀树中插入字符串 word 。
	boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
	boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

插入字符串
	我们从字典树的根开始，插入字符串。对于当前字符对应的子节点，有两种情况：
		子节点存在。沿着指针移动到子节点，继续处理下一个字符。
		子节点不存在。创建一个新的子节点，记录在 \textit{children}children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
	重复以上步骤，直到处理字符串的最后一个字符，然后将当前节点标记为字符串的结尾。
查找前缀
	我们从字典树的根开始，查找前缀。对于当前字符对应的子节点，有两种情况：
		子节点存在。沿着指针移动到子节点，继续搜索下一个字符。
		子节点不存在。说明字典树中不包含该前缀，返回空指针。
	重复以上步骤，直到返回空指针或搜索完前缀的最后一个字符。
	若搜索到了前缀的末尾，就说明字典树中存在该前缀。此外，若前缀末尾对应节点的 \textit{isEnd}isEnd 为真，则说明字典树中存在该字符串。

*/
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if this.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
