package main

import (
	"fmt"
	"math"
)

/*
请⽤简单可扩展的代码，实现32位有符号整数的⼆元运算（不考虑整数溢出），⽀持如下运算：
assert(BinaryOp(1, "+", 1, &r) == 0);
assert(r == 2);
assert(BinaryOp(2, "**", 2, &r) == 0);
assert(r == 4);
assert(BinaryOp(1, "/", 0, &r) != 0); // div by zero
*/

func BinaryOp(a int, op string, b int, ptr *int) (res int) {
	switch op {
	case "+":
		*ptr = a + b
		return
	case "-":
		*ptr = a - b
		return
	case "*":
		*ptr = a * b
		return
	case "/":
		if b == 0 {
			return 1
		}
		*ptr = a / b
		return
	case "**":
		if b == 0 {
			return 1
		}
		*ptr = int(math.Pow(float64(a), float64(b)))
		return
	case "%":
		if b == 0 {
			return 1
		}
		*ptr = a % b
		return
	default:
		return 1
	}
}

func main() {
	var r int
	fmt.Println(BinaryOp(1, "+", 1, &r))
	fmt.Println(r == 2)
	fmt.Println(BinaryOp(2, "**", 2, &r))
	fmt.Println(r == 4)
	fmt.Println(BinaryOp(1, "/", 0, &r))
}
