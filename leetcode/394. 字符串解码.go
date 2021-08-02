package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。

如果当前的字符为数位，解析出一个数字（连续的多个数位）并进栈
如果当前的字符为字母或者左括号，直接进栈
如果当前的字符为右括号，开始出栈，一直到左括号出栈，出栈序列反转后拼接成一个字符串，此时取出栈顶的数字（此时栈顶一定是数字，想想为什么？），就是这个字符串应该出现的次数，我们根据这个次数和字符串构造出新的字符串并进栈

*/
func decodeString(s string) string {
	stack := []string{}
	for _, i := range s {
		if i >= '0' && i <= '9' || i == '[' || (i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z') {
			stack = append(stack, string(i))
		} else {
			sub := []string{}
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			reverseSlice(sub)
			stack = stack[:len(stack)-1]
			num := []string{}
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				num = append(num, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			reverseSlice(num)
			rep, _ := strconv.Atoi(getString(num))
			t := strings.Repeat(getString(sub), rep)
			stack = append(stack, t)
		}
	}
	return getString(stack)
}
func getString(s []string) string {
	re := ""
	for _, i := range s {
		re += i
	}
	return re
}
func reverseSlice(s []string) {
	for j := 0; j < len(s)/2; j++ {
		s[j], s[len(s)-j-1] = s[len(s)-j-1], s[j]
	}
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}
