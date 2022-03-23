package main

import (
"fmt"
"strconv"
)

var stack1 []string
var stack2 []int

func calculate(m string) int {
	var numCnt int
	for _,s := range m {
		i := string(s)
		i2, err := strconv.Atoi(i)
		// 判断左括号
		if i == "(" {
			stack1 = append(stack1, i)
			numCnt = 0
		} else if i == ")" {
			// 判断右括号，将左括号以来的操作符全部计算
			if stack1[len(stack1)-1] == "-" && numCnt == 1 {
				// 负数
				num := - stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				stack1 = stack1[:len(stack1)-2]
				PutNumInStack(num)
			} else {
				// 计算两个括号之间的结果
				calculateResult()
			}
			numCnt = 0
		} else if err == nil {
			//判断数字
			numCnt++
			PutNumInStack(i2)
		} else {
			stack1 = append(stack1, i)
			// //判读操作符
			// if i == "*" || i == "/" {
			//     // 乘除法先入栈
			// }
		}
	}
	calculateResult()
	if len(stack2) == 1 {
		return stack2[0]
	}
	panic("calculate error")
}

func PutNumInStack(num int) {
	stack2 = append(stack2, num)
	if len(stack1) > 0 {
		lastOp := stack1[len(stack1)-1]
		if lastOp == "*" || lastOp == "/" {
			stack1 = stack1[:len(stack1)-1]
			num1,num2 := stack2[len(stack2)-2],stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-2]
			if lastOp == "*" {
				stack2 = append(stack2, num1 * num2)
			} else {
				stack2 = append(stack2, num1 / num2)
			}
		}
	}

}

func calculateResult() {
	for len(stack1) > 0 {
		lastOp := stack1[len(stack1)-1]
		if lastOp != "(" {
			num1,num2 := stack2[len(stack2)-2],stack2[len(stack2)-1]
			stack1 = stack1[:len(stack1)-1]
			stack2 = stack2[:len(stack2)-2]
			if lastOp == "-" {
				PutNumInStack(num1-num2)
			} else if lastOp == "+" {
				PutNumInStack(num1+num2)
			}
		} else {
			stack1 = stack1[:len(stack1)-1]
			break
		}
	}
}

func main() {
	fmt.Println(calculate("(1+9)*3+3*(-4)"))
}
