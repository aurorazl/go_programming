package main

import "fmt"

/*
请根据每日 气温 列表 temperatures，重新生成一个列表，要求其对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
如果气温在这之后都不会升高，请在该位置用0 来代替。

单调递减栈

*/

//逆序
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	ret := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if len(stack) == 0 { // 初始化
			stack = append(stack, i)
		} else if temperatures[stack[len(stack)-1]] > temperatures[i] { //比栈顶小
			ret[i] = stack[len(stack)-1] - i
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] { //比栈顶大，需要替换
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				ret[i] = stack[len(stack)-1] - i
			}
			stack = append(stack, i)
		}
	}
	return ret
}

//正序
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] { // 不断更新之前的还没到高的
			ret[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}

func main() {
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
