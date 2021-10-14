package main

import "fmt"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
x→x^2→x^4→x^8→x^16 →x^32 →x^64
*/

func myPow(x float64, n int) float64 {
	ret := quickPow(x, n)
	if n < 0 {
		ret = 1 / ret
	}
	return ret
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
func main() {
	fmt.Print(myPow(2, -1))
}
