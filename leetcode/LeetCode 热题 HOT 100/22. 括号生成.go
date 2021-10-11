package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
广度搜索思路：
	字典结构：
		当前字符串：[左括号剩余未匹配数量x，剩余可用左括号次数y，剩余可用右括号次数z]
		优化：【剩余可用左括号次数y，剩余可用右括号次数z】，而左括号剩余未匹配数量 x 只需要判断 y < z 即可。
			  不需要两个map，只需要一个队列，队列不断取出和放入。满足条件的放入res数组。最终返回即可。
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

type Node struct {
	res       string
	left_num  int
	right_num int
}

// 广度优化优化版
func generateParenthesisAdvanted(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	queue := []*Node{{"(", n - 1, n}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.left_num == 0 && node.right_num == 0 {
			res = append(res, node.res)
		}
		if node.left_num > 0 {
			queue = append(queue, &Node{node.res + "(", node.left_num - 1, node.right_num})
		}
		if node.right_num > 0 && node.left_num < node.right_num {
			queue = append(queue, &Node{node.res + ")", node.left_num, node.right_num - 1})
		}
	}
	return res
}

//深度优化，利用系统栈进行回溯
func generateParenthesisByDepth(n int) []string {
	res := []string{}
	var dsp func(curStr string, left, right int)
	dsp = func(curStr string, left, right int) {
		if left == 0 && right == 0 {
			res = append(res, curStr)
		}
		if left > 0 {
			dsp(curStr+"(", left-1, right)
		}
		if right > 0 && left < right {
			dsp(curStr+")", left, right-1)
		}
	}
	dsp("(", n-1, n)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
