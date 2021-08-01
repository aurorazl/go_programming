package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
解析：
	栈
*/
func isValid(s string) bool {
	strs := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}
	for _, i := range s {
		if _, ok := strs[string(i)]; !ok {
			stack = append(stack, string(i))
		} else {
			if len(stack) == 0 || strs[string(i)] != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()(**"))
}
