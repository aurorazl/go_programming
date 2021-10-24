package main

import (
	"fmt"
	"math/rand"
)

/*
设计一个支持在平均时间复杂度 O(1)下，执行以下操作的数据结构：

insert(val)：当元素 val 不存在时返回 true，并向集合中插入该项，否则返回 false 。
remove(val)：当元素 val 存在时返回 true，并从集合中移除该项，否则返回 false。
getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。

哈希表（存储数组下标） + 数组（等概率返回）
*/

type RandomizedSet struct {
	dict map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}
	this.list = append(this.list, val)
	this.dict[val] = len(this.list) - 1
	return true
}

//数组中后面的元素会向前移动，时间复杂度为O(n)。交换位置可以节省时间。再更新那个元素在字典的索引，删除最后一个元素
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.dict[val]; ok {
		this.list[index], this.list[len(this.list)-1] = this.list[len(this.list)-1], this.list[index]
		this.dict[this.list[index]] = index
		delete(this.dict, val)
		this.list = this.list[:len(this.list)-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := len(this.list)
	return this.list[rand.Intn(n)]
}

func main() {
	randomizedSet := Constructor()
	randomizedSet.Insert(0)
	randomizedSet.Insert(1)
	fmt.Println(randomizedSet.Remove(0))
	fmt.Println(randomizedSet.Insert(2))
	fmt.Println(randomizedSet.Remove(1))
	fmt.Println(randomizedSet.GetRandom())
}
