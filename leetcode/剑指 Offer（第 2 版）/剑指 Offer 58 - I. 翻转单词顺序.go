package main

import "fmt"

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

栈法
*/

func reverseWords(s string) string {
	s = s + " " //方便兼容读取最后一个单词
	start := 0
	stack := []string{}
	for i := range s {
		if s[i] == ' ' {
			stack = append(stack, s[start:i])
			start = i + 1
		}
		i++
	}
	ret := ""
	for i := len(stack) - 1; i >= 0; i-- { // 倒序输出
		if stack[i] == "" {
			continue
		}
		ret += stack[i]
		ret += " "
	}
	if len(ret) > 0 {
		return ret[:len(ret)-1] //根据情况忽略最后一个空格
	}
	return ret
}
func main() {
	fmt.Println(reverseWords(" "))
}
