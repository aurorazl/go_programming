package main

import "fmt"

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
遍历递归，求出每种可能的长度进行比较即可。
思路：
2.遍历给字符串中的所有字符
    2.1若当前字符s[index]为左括号'('，将当前字符下标index入栈（下标稍后有其他用处），处理下一字符
    2.2若当前字符s[index]为右括号')'，判断当前栈是否为空
        2.2.1若栈为空，则start = index + 1，处理下一字符（当前字符右括号下标不入栈）
        2.2.2若栈不为空，则出栈（由于仅左括号入栈，则出栈元素对应的字符一定为左括号，可与当前字符右括号配对），判断栈是否为空
            2.2.2.1若栈为空，则max = max(max, index-start+1)
            2.2.2.2若栈不为空，则max = max(max, index-栈顶元素值)
*/

func longestValidParentheses(s string) int {
	stack := []int{}
	start := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				start = i + 1
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					maxLen = max(maxLen, i-start+1)
				} else {
					maxLen = max(maxLen, i-stack[len(stack)-1])
				}
			}
		}
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("(()()"))
}
