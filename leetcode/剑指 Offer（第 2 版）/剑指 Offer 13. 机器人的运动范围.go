package main

import "fmt"

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

暴力解法，时间复杂度为O（mn）
深度或者广度优先搜索
*/

func movingCount(m int, n int, k int) int {
	moves := [2][2]int{{1, 0}, {0, 1}}
	ans := 0
	queue := [][2]int{{0, 0}}
	visit := make([][]int, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] < m && cur[1] < n && get(cur[0])+get(cur[1]) <= k && visit[cur[0]][cur[1]] == 0 {
			ans++
			visit[cur[0]][cur[1]] = 1
			for _, move := range moves {
				queue = append(queue, [2]int{cur[0] + move[0], cur[1] + move[1]})
			}
		}
	}
	return ans
}
func get(num int) int {
	ret := 0
	for num > 0 {
		ret += num % 10
		num /= 10
	}
	return ret
}
func main() {
	fmt.Println(movingCount(3, 2, 17))
}
