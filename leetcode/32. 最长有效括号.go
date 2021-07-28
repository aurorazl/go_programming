package main

import "fmt"

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
遍历递归，求出每种可能的长度进行比较即可。
思路：
	遍历字符串
    如果是（，
  		如果之前栈深度>0，继续遍历
		如果深度=0，继续遍历
	如果是），
		如果之前栈深度>1，length+=2，更新最大值
		如果深度=1，将length保存到list，length=0
		如果之前栈深度=0，说明中断了，跳过当前符号，合并list，更新最大值，清空list
*/

func longestValidParentheses(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			curLen := calculateLen(s[i:])
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}

func calculateLen(s string) int {
	stack := 0
	length := 0
	tmp := []int{}
	maxLen := 0
	for _, i := range s {
		if string(i) == "(" {
			if stack > 0 {
				tmp = append(tmp, length)
				length = 0
			}
			stack++

		} else {
			if stack >= 1 {
				length += 2
				if length > maxLen {
					maxLen = length
				}
				if stack == 1 {
					sumLen := sum(tmp) + length
					if sumLen > maxLen {
						maxLen = sumLen
					}
					tmp = []int{sumLen}
					length = 0
				}
				stack--
			} else {
				sumLen := sum(tmp)
				if sumLen > maxLen {
					maxLen = sumLen
				}
				tmp = []int{}
			}
		}
	}
	if len(tmp) != 0 && stack == 0 {
		sumLen := sum(tmp)
		if sumLen > maxLen {
			maxLen = sumLen
		}
		tmp = []int{}
	}
	return maxLen
}
func sum(list []int) int {
	val := 0
	for _, i := range list {
		val += i
	}
	return val
}
func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
