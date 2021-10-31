package main

/*
给定一个由 0 和 1 组成的矩阵 mat，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。

广度优先搜索，每个0的下标先入队，然后遍历队列每个下标的上下左右,取非0的，再次入队，更新距离
深度优先，非0的下标先入队，然后每个下标不断的遍历多次上下左右，直到找到0，更新距离。
*/

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	for len(queue) != 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			nx := x + dx[k]
			ny := y + dy[k]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && mat[nx][ny] == 1 {
				if res[nx][ny] == 0 || res[x][y]+1 < res[nx][ny] {
					res[nx][ny] = res[x][y] + 1
					queue = append(queue, [2]int{nx, ny}) // 希望再次更新连接的其他下标
				}
			}
		}
	}
	return res
}
