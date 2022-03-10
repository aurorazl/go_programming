package main


import (
	"fmt"
	"math"
)
type TreeNode struct{
	val int
	left *TreeNode
	right *TreeNode
}

func getLine(x1,y1,x2,y2 int)(k,b int){
	k = (y1-y2)/(x1-x2)
	b = y1-x1*k
	return
}
func PointInLineCnt(list [][]int,k int,b int) int {
	cnt := 0
	for _,one := range list{
		if (one[0]*k + b)==one[1]{
			cnt ++
		}
	}
	return cnt
}

func FindMostPointLine(list [][]int)(int,int){
	cnt := 0
	x := 0
	y := 0
	for i:=0;i<len(list);i++ {
		for j:=i+1;j<len(list);j++{
			k,b := getLine(list[i][0],list[i][1],list[j][0],list[j][1])
			if PointInLineCnt(list,k,b)>cnt{
				x = k
				y = b
			}
		}
	}
	return x,y
}

func FindMostPoint(list [][]int) int {
	ans := 0
	//x := 0.0
	//y := 0.0
	if len(list)<=2{
		return len(list)
	}
	for i:=0;i<len(list);i++ {
		for j:=i+1;j<len(list);j++{
			cnt := 2
			for k:=j+1;k<len(list);k++{
				if list[i][0]-list[j][0] == 0 && list[j][0]-list[k][0]==0 {
					cnt++
					continue
				}
				if list[i][0]-list[j][0] == 0 ||  list[j][0]-list[k][0]==0 {
					continue
				}
				//x = float64(list[i][1]-list[j][1]) / float64(list[i][0]-list[j][0])
				//y = float64(list[i][1]) - x * float64(list[j][0])
				if float64(list[i][1]-list[j][1]) / float64(list[i][0]-list[j][0]) == float64(list[j][1]-list[k][1])/ float64(list[j][0]-list[k][0]){
					cnt++
				}
			}
			ans = int(math.Max(float64(ans), float64(cnt)))
		}
	}
	return ans
}

func main() {
	list := [][]int{{0,0},{4,5},{7,8},{8,9},{5,6},{3,4},{1,1}}
	c := FindMostPoint(list)
	fmt.Println(c)


}
