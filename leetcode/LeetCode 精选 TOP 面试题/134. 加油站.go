package main

/*
在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

思路：
从头到尾遍历每个加油站，并检查以该加油站为起点，最终能否行驶一周。我们可以通过减小被检查的加油站数目，来降低总的时间复杂度。
如何降低？
假设我们此前发现，从加油站 xx 出发，每经过一个加油站就加一次油（包括起始加油站），最后一个可以到达的加油站是 y（不妨设 x<y）。
从 x,y 之间的任何一个加油站出发，都无法到达加油站 y 的下一个加油站。
那么：
	我们首先检查第 0 个加油站，并试图判断能否环绕一周；如果不能，就从第一个无法到达的加油站开始继续检查。
*/

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; {
		cur := 0
		sumGas, sumCost := 0, 0
		for cur < n {
			start := (i + cur) % n
			sumGas += gas[start]
			sumCost += cost[start]
			if sumCost > sumGas {
				break
			}
			cur++
		}
		if cur == n {
			return i
		} else {
			i += cur + 1
		}
	}
	return -1
}
