package main

import (
	"fmt"
	"sort"
)

type Range struct {
	From int
	To   int
}
type RangeList []Range

func (rangeList RangeList) Len() int {
	return len(rangeList)
}
func (rangeList RangeList) Less(i int, j int) bool {
	return rangeList[i].From < rangeList[j].From
}
func (rangeList RangeList) Swap(i int, j int) {
	rangeList[i], rangeList[j] = rangeList[j], rangeList[i]
}

func mergeLi(rangeList RangeList) []Range {
	sort.Sort(rangeList)
	tmp := RangeList{}
	for i := 0; i < len(rangeList); i += 2 {
		if rangeList[i].To == rangeList[i+1].From-1 {
			tmp = append(tmp, Range{From: rangeList[i].From, To: rangeList[i+1].To})
		} else {
			tmp = append(tmp, rangeList[i:i+2]...)
		}
	}
	return tmp
}

func main() {
	a := RangeList{
		Range{4, 5},
		Range{1, 2},
	}
	fmt.Println(mergeLi(a))
}
