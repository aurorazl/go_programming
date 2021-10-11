package LeetCode_热题_HOT_100

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/
func numIslands(grid [][]byte) int {
	nums_island := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				nums_island++
				dfs(grid, i, j)
			}
		}
	}
	return nums_island
}
func dfs(grid [][]byte, i, j int) { //维护一个栈也可以
	grid[i][j] = '0'
	for _, one := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
		new_i := i + one[0]
		new_j := j + one[1]
		if new_i >= 0 && new_i < len(grid) && new_j >= 0 && new_j < len(grid[0]) && grid[new_i][new_j] == '1' {
			dfs(grid, new_i, new_j)
		}
	}
}
