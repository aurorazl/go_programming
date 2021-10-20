package main

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。

比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

切片
辅助数组，两次存储两部分
字符串遍历拼接（res += s[i%len()] for i in k,len()+k）。就是先遍历后半部分，在遍历前半部分。
*/

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWords3(s string, n int) string {
	ret := ""
	for i := n; i < len(s)+n; i++ {
		ret += string(s[i%len(s)])
	}
	return ret
}
