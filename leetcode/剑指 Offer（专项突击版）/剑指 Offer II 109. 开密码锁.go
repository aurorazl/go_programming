package main

import "fmt"

/*

一个密码锁由 4 个环形拨轮组成，每个拨轮都有 10 个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，请给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

广度优先搜索

在一开始将 (0000, 0)(0000,0) 加入队列，并使用该队列进行广度优先搜索。
在搜索的过程中，设当前搜索到的数字为 status，旋转的次数为 step，我们可以枚举 status 通过一次旋转得到的数字。
设其中的某个数字为 next_status，如果其没有被搜索过，我们就将 (next_status,step+1) 加入队列。
如果搜索到了 target，我们就返回其对应的旋转次数。
*/
type pair struct {
	status string
	step   int
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return -1
	}
	queue := []pair{pair{"0000", 0}}
	dead := make(map[string]bool, len(deadends))
	for _, deadend := range deadends {
		dead[deadend] = true
	}
	if dead["0000"] {
		return -1
	}
	seen := make(map[string]bool)
	for len(queue) != 0 { // 层次遍历，会越来越深
		one := queue[0]
		queue = queue[1:]
		for _, next := range spin(one) {
			if !seen[next] && !dead[next] {
				if next == target {
					return one.step + 1
				}
				seen[next] = true
				queue = append(queue, pair{next, one.step + 1})
			}
		}
	}
	return -1
}

func spin(one pair) []string {
	res := []string{}
	s := []byte(one.status)
	for i, char := range s {
		s[i] = char - 1
		if s[i] < '0' {
			s[i] = '9'
		}
		res = append(res, string(s))
		s[i] = char + 1
		if s[i] > '9' {
			s[i] = '0'
		}
		res = append(res, string(s))
		s[i] = char
	}
	return res
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
