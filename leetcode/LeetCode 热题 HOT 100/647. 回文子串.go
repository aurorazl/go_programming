package LeetCode_热题_HOT_100

import "fmt"

/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

由此我们可以看出长度为 n 的字符串会生成 2n-1 组回文中心
这样我们只要从 0 到 2n−2 遍历 i，就可以得到所有可能的回文中心，这样就把奇数长度和偶数长度两种情况统一起来了。

*/
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
func main() {
	fmt.Print(countSubstrings("abc"))
}
