package main

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

方法一：使用哈希表存储频数+两次遍历
方法二：使用哈希表存储索引
	两次的频数改为-1，遍历一次哈希映射中的所有值，找出其中不为 -1−1 的最小值
方法三：队列

*/
