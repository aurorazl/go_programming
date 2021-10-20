package main

/*
0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

圆圈可以通过 % 除余实现。
class Solution {
    public int lastRemaining(int n, int m) {
        return f(n, m);
    }
    public int f(int n, int m) {
        if (n == 1) {
            return 0;
        }
        int x = f(n - 1, m);
        return (m + x) % n;
    }
}
*/

func lastRemaining(n int, m int) int {
	return f(n, m)
}
func f(n, m int) int {
	if n == 1 {
		return 0 // 最后只有一个元素时，幸存下标0
	}
	x := f(n-1, m)
	return (m + x) % n // 前面遗留的元素，在上一轮的位置肯定是第m+1位
}

//方法二：数学 + 迭代
