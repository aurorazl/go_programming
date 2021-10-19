package main

import "fmt"

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

考虑借用一个辅助栈 stack，模拟 压入 / 弹出操作的排列。根据是否模拟成功，即可得到结果。
入栈操作： 按照压栈序列的顺序执行。
出栈操作： 每次入栈后，循环判断 “栈顶元素 == 弹出序列的当前元素” 是否成立，将符合弹出序列顺序的栈顶元素全部弹出。
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{} // 辅助栈，用来存放哪些不能及时弹出的元素
	for _, onePushed := range pushed {
		stack = append(stack, onePushed) // 每个元素都先放入辅助栈，再决定要不要取出来
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
