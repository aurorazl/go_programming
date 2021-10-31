package main

import "fmt"

/*
给定一个有个节点的有向无环图，用二维数组graph表示，请找到所有从0到n-1的路径并输出（不要求按顺序）。
graph的第 i 个数组中的单元都表示有向图中 i号节点所能到达的下一些结点（译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ），
若为空，就是没有下一个节点了。

深度优先搜索(回溯)
*/

func allPathsSourceTarget(graph [][]int) (res [][]int) {
	var dfs func(int)
	path := []int{0}
	dfs = func(x int) {
		if x == len(graph)-1 {
			res = append(res, append([]int{}, path...))
		}
		for _, y := range graph[x] {
			path = append(path, y)
			dfs(y)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
