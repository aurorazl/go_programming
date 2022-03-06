package main

import "fmt"
import "time"

type Point struct {
	x, y int
}

type a *int

func (p *Point) ScaleBy(bei int) int {
	return p.x * bei
}
func (p Point) ScaleBy2(bei int) {

}

func (p *a) ScaleBy3(bei int) {

}

func main() {
	fmt.Print((&Point{1, 2}).ScaleBy(2))
	a := &Point{1, 2}
	a.ScaleBy2(2)
	b := Point{1, 2}
	b.ScaleBy(2 * int(2))
}
