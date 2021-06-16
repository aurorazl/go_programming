package main

import "fmt"

type student struct{ Name string }

func (st student) Name1() {

}
func f1(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Printf("%T", msg)
	}
}

type Integer int

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy2(factor float64) {
	fmt.Println("ss")
}

func main() {
	f1(student{"s"})

	var a Integer = 1
	//var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer)
	fmt.Printf("%T", sum)

	(Point{1, 2}).ScaleBy2(2)

	var b = make([]int, 0)
}
