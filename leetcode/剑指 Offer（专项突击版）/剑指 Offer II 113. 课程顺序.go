package main

/*
现在总共有 numCourses门课需要选，记为0到numCourses-1。
给定一个数组prerequisites ，它的每一个元素prerequisites[i]表示两门课程之间的先修顺序。例如prerequisites[i] = [ai, bi]表示想要学习课程 ai，需要先完成课程 bi。
请根据给出的总课程数 numCourses 和表示先修顺序的prerequisites得出一个可行的修课序列。
可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

拓扑排序
在拓扑排序中有两个重要概念，即出度和入度。
一种常用的拓扑排序算法是每次都从有向无环图中取出入度为 0 的节点添加到拓扑排序的序列中，然后删除以该节点为起点的边。
重复以上过程，直至图为空或者不存在入度为 0 的节点，若最终图为空，那么图就是一个有向无环图，若最终图不为空且已不存在入度为 0 的节点，那么该图一定有环。

在遍历节点时采用广度优先搜素算法，使用队列保存入度为 0 的节点。
每次从队列取出一个节点，就把该节点添加到排序序列中，同时将该先修课对应的后修课节点的入度减 1，这相当于删除从先修课到后修课的边。
最终若拓扑排序序列的节点数等于总课程数，则存在一个可行的拓扑排序，反之则图内必有环，则不存在拓扑排序。

*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	indo := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, course := range prerequisites {
		ai, bi := course[0], course[1]
		graph[bi] = append(graph[bi], ai)
		indo[ai]++
	}
	queue := []int{} //学习顺序
	for ai, cnt := range indo {
		if cnt == 0 {
			queue = append(queue, ai) //先取出入度为0的
		}
	}
	for i := 0; i < len(queue); i++ {
		bi := queue[i]
		for _, ai := range graph[bi] {
			indo[ai]--
			if indo[ai] == 0 {
				queue = append(queue, ai)
			}
		}
	}
	if len(queue) < numCourses {
		return []int{}
	}
	return queue
}
