package main

/*
在字典（单词列表）wordList 中，从单词 beginWord和 endWord 的 转换序列 是一个按下述规格形成的序列：

序列中第一个单词是 beginWord 。
序列中最后一个单词是 endWord 。
每次转换只能改变一个字母。
转换过程中的中间单词必须是字典wordList 中的单词。
给定两个长度相同但内容不同的单词 beginWord和 endWord 和一个字典 wordList ，找到从beginWord 到endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。

广度优先搜索 + 优化建图
本题要求的是最短转换序列的长度，看到最短首先想到的就是广度优先搜索。
想到广度优先搜索自然而然的就能想到图，但是本题并没有直截了当的给出图的模型，因此我们需要把它抽象成图的模型。

我们可以把每个单词都抽象为一个点，如果两个单词可以只改变一个字母进行转换，那么说明他们之间有一条双向边。
因此我们只需要把满足转换条件的点相连，就形成了一张图。

*/
