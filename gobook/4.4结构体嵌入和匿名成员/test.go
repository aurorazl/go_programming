package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      *string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID []int
}

var struct_pointer *Employee

func EmployeeByID(id int) Employee { return Employee{} }

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}
func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}
func main() {
	fmt.Print(EmployeeByID(0).Salary)
	//EmployeeByID(0).Salary = 0
	struct_pointer = &Employee{}
	*struct_pointer = Employee{}
	//fmt.Println(Employee{}==Employee{})
	fmt.Println(namedReturnValues())
}
