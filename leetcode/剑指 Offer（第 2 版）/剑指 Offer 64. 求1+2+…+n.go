package main

/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。


遍历
快速乘 -- 位运算
*/

func sumNums(n int) (ret int) {
	for i := 1; i <= n; i++ {
		ret += i
	}
	return
}
