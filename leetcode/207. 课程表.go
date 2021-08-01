package main

/*
你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。
例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, one := range prerequisites {
		edge[one[1]] = append(edge[one[1]], one[0])
		indeg[one[0]]++
	}
	q := []int{}
	for index, i := range indeg {
		if i == 0 {
			q = append(q, index)
		}
	}
	res := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		res = append(res, u)
		for _, i := range edge[u] {
			indeg[i]--
			if indeg[i] == 0 {
				q = append(q, i)
			}
		}
	}
	return len(res) == numCourses
}
