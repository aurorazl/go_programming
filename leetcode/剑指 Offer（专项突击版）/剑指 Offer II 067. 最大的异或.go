package main

import "fmt"

/*
给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。

1.构建前缀树：将所有的数都插入前缀树中。与普通前缀树不同，这里是根据一个数的二进制位来插入，普通前缀树是根据单词中的字符来插入。
还有一点需要注意的是，我们优先考虑高位，所以是从高位到低位进行插入。(每个单词补位够30位)
2.对于 nums 中的每个数 num ，都到前缀树中去搜索num最大异或的那个数，然后计算最大异或值，最后，从这些异或值中挑出最大的一个就是要的答案。
重点：搜索的方法
	异或值最大，我们就要尽量让每个异或位都和 num 对应的二进制位不同。
	如果 num 当前位为 0，就到 next[1] 去搜索；
	如果 num 当前位为 1，就到 next[0] 去搜索;
	如果与 num 当前位相反的那一位为空，那就只能到相同的那一位去搜索了。

*/

type Trie struct {
	children [2]*Trie
	isNode   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(num int) {
	node := this
	for i := 30; i >= 0; i-- { //每个单词都有30位二进制码
		d := num >> i & 1
		trie := node.children[d]
		if trie != nil {
			node = trie
		} else {
			node.children[d] = &Trie{}
			node = node.children[d]
		}
	}
}

func (this *Trie) Search(num int) int {
	node := this
	xorNum := 0
	for i := 30; i >= 0; i-- {
		d := num >> i & 1
		other := (d - 1) * -1
		if node.children[other] != nil {
			node = node.children[other]
			xorNum = xorNum*2 + 1
		} else {
			node = node.children[d]
			xorNum = xorNum * 2
		}
	}
	return xorNum
}

func findMaximumXOR(nums []int) int {
	trie := Constructor()
	for _, num := range nums {
		trie.Insert(num)
	}
	ret := 0
	for _, num := range nums {
		ret = max(trie.Search(num), ret)
	}
	return ret
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}
