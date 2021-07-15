package main

import "fmt"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 */

func findMaxSum(list []int)int {
	if len(list)==0{
		return 0
	}
	maxSum := list[0]
	sum := list[0]
	for _,i := range list[1:]{
		if sum < 0 {
			sum = i
		}else{
			if sum + i > 0{
				sum += i
			}else{
				sum = i
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}


func main(){
	tp := []int{-8,-3,2,-4,3,-2,-1,-9}
	fmt.Print(findMaxSum(tp))
}
20

19 18 16
17
15

14 16

1  2 4
1

14 确保

10 - 14 

6 -

2

