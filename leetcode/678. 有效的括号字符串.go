package main

import "fmt"

/*
延伸，字符串包含（，*，），*可以当作（）或者空
栈(先进后出)
定义两个栈，leftStack 存 ' ( ' 所在位置的下标，starStack 存 '*' 所在位置的下标。
1.当遇到 ' ( ' 时，' ( ' 所在位置的下标入栈；当遇到 ' * ' 时，' * ' 所在位置的下标入栈。
2.当遇到 ' ) ' 时，要令 leftStack 中的栈顶元素出栈，若此时 leftStack 为空，则继续看 starStack 是否为空，若不为空则 starStack 栈顶元素出栈，若为空返回则 false (遇到了 ' ) '，但在这之前一个 ' ( ' 和 ' * ' 都没遇到，则一定不会匹配)。
3.当字符串全部遍历完时，若 starStack 的长度比 leftStack 的长度要小，则返回 false (有剩余的 ' ( '，但 ' * ' 的数量不够了，
则一定不会匹配)；
否则，比较两个栈的栈顶元素值的大小，要保证 ' ( ' 在 ' * ' 的左边(starStack.peek() > leftStack.peek())才能匹配成功，
当遇到满足条件的栈顶元素时，栈顶元素出栈，继续比较下一个。只要有一次该条件不满足，则直接返回 false；否则，返回 true。
证明：
	先扣左括号，
	如果剩下的还是左括号在前，不匹配
	如果在后，那么根据消消乐，肯定可以消掉
*/
func isValid(s string) bool {
	left_stack := []int{}
	star_stack := []int{}
	for index, i := range s {
		if i == '(' {
			left_stack = append(left_stack, index)
		} else if i == '*' {
			star_stack = append(star_stack, index)
		} else {
			if len(left_stack) > 0 {
				left_stack = left_stack[:len(left_stack)-1] // 对于一个星，应该先去掉左括号，
			} else if len(star_stack) > 0 {
				star_stack = star_stack[:len(star_stack)-1]
			} else {
				return false
			}
		}
	}
	if len(star_stack) >= len(left_stack) {
		for len(left_stack) > 0 {
			if left_stack[len(left_stack)-1] > star_stack[len(star_stack)-1] {
				return false
			}
			left_stack = left_stack[:len(left_stack)-1]
			star_stack = star_stack[:len(star_stack)-1]
		}
		return true
	}
	return false
}
func main() {
	fmt.Println(isValid("(*)(*"))
}
