package main

/*
输入list[["A","B"],["A","C"],["B","D"],["D","A"]]，判断是否闭环

113.课程顺序
 */

type Node struct {
	val string
	subNode map[string]*Node
}

func newNodeInstance(val string) *Node {
	return &Node{val: val, subNode: make(map[string]*Node)}
}

func isCircle(list [][]string) bool {
	// 一个服务被依赖的所有服务
	graph := make(map[string][]string)

	//一个服务的依赖次数
	indo := make(map[string]int)

	for _, one := range list {
		dependencyPar := one[0]
		dependencySub := one[1]
		graph[dependencySub] = append(graph[dependencySub], dependencyPar)
		indo[dependencyPar] ++
		indo[dependencySub] = 0
	}

	queue := []string{}

	for d,cnt := range indo {
		if cnt == 0 {
			queue = append(queue, d)
		}
	}

	for i := 0; i < len(queue); i++ {
		dependencySub := queue[i]
		for _,dependencyPar := range graph[dependencySub] {
			indo[dependencyPar]--
			if indo[dependencyPar] == 0 {
				queue = append(queue, dependencyPar)
			}
		}
	}
	for _, cnt := range indo {
		if cnt != 0 {
			return true
		}
	}
	return false
}

func main() {
	println(isCircle([][]string{{"A", "B"}, {"A", "C"}, {"B", "D"}, {"D", "E"}, {"E","B"}}))
}