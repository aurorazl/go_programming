package main

import "fmt"

/*
给定一个整数数组 asteroids，表示在同一行的小行星。
对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。

只有一种情况会导致小行星发生碰撞，那就是当前小行星往左飞行，而上一个小行星往右飞行时。
栈
*/

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, num := range asteroids {
		for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1]+num < 0 { // a+b<0，两个肯定有正有负，而且负的大
			stack = stack[:len(stack)-1] //1.不断碰撞摧毁之前的
		}
		if len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] == -num {
			stack = stack[:len(stack)-1] // 2.两个都爆炸
		} else if len(stack) == 0 || stack[len(stack)-1] < 0 || num > 0 { //3.入栈
			stack = append(stack, num)
		} //4.被摧毁
	}
	return stack
}
func main() {
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
}
