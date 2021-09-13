package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
广度搜索思路：
	字典结构：当前字符串：[左括号剩余未匹配数量x，剩余可用左括号次数y，剩余可用右括号次数z]
    判断逻辑
    1.如果长度x>=1
        1.1 如果y>=1,可以选择左括号、右括号，遍历两种可能
        1.2 如果y==0，接下来只能右括号，x--
    2.如果长度x=0，只能左括号
        1.1 y>=1，只能左括号
        1.2 y==0，结束退出
深度优先：
    栈调用
*/

//广度
func generateParenthesis(n int) []string {
	res := map[string][]int{"(": []int{1, n - 1, n}}
	for {
		cur_len := 0
		tmp := map[string][]int{}
		for str, state := range res {
			if state[0] >= 1 {
				if state[1] >= 1 {
					tmp[str+"("] = []int{state[0] + 1, state[1] - 1, state[2]}
					tmp[str+")"] = []int{state[0] - 1, state[1], state[2] - 1}
				} else {
					tmp[str+")"] = []int{state[0] - 1, state[1], state[2] - 1}
				}
			} else {
				if state[1] >= 1 {
					tmp[str+"("] = []int{state[0] + 1, state[1] - 1, state[2]}
				} else {
					cur_len++
				}
			}
		}
		if cur_len == len(res) {
			break
		} else {
			res = tmp
		}
	}
	final := []string{}
	for k, v := range res {
		if v[2] != 0 {
			for i := 0; i < v[2]; i++ {
				k += ")"
			}
		}
		final = append(final, k)
	}
	return final
}
func main() {
	fmt.Println(generateParenthesis(3))
}
