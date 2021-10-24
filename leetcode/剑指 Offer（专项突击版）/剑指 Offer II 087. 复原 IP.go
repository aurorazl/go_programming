package main

import (
	"fmt"
	"strconv"
)

/*
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
*/

// 回溯法
func restoreIpAddresses(s string) (ans []string) {
	n := len(s)
	var dfs func(int, string, int)
	dfs = func(left int, str string, cnt int) {
		if cnt > 4 {
			return
		}
		if cnt == 4 && left == n {
			ans = append(ans, str[:len(str)-1])
		}
		for i := 0; i < 3 && left+i < n; i++ {
			num, _ := strconv.Atoi(s[left : left+i+1])
			if s[left] == '0' {
				dfs(left+1, str+"0.", cnt+1)
				break
			}
			if num >= 1 && num <= 255 {
				dfs(left+i+1, str+s[left:left+i+1]+".", cnt+1)
			}
		}
	}
	dfs(0, "", 0)
	return
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
