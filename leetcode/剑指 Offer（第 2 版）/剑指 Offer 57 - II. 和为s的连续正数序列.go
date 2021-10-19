package main

import "fmt"

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

双指针形成滑动窗口
起始 l=1,r=2。
如果 sum<target 则说明指针 r 还可以向右拓展使得 sum 增大，此时指针 r 向右移动，即 r+=1
如果 >target 则说明以 l 为起点不存在一个 r 使得 sum=target ，此时要枚举下一个起点，指针 l 向右移动，即l+=1
如果 sum==target 则说明我们找到了以 l 为起点得合法解 [l,r] ，我们需要将 [l,r] 的序列放进答案数组，且我们知道以 l为起点的合法解最多只有一个，
所以需要枚举下一个起点，指针 l 向右移动，即 l+=1

*/

func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	for l, r := 1, 2; l < r; {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			perm := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				perm[i-l] = i
			}
			ans = append(ans, perm)
			l++
		} else if sum > target {
			l++
		} else {
			r++
		}
	}
	return ans
}
func main() {
	fmt.Println(findContinuousSequence(9))
}
