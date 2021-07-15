package main

import (
	"fmt"
)

// Animal 动物基类
type Animal struct {
	name string
}

func (a *Animal) Play() {
	fmt.Println(a.Speak())
}

func (a *Animal) Speak() string {
	return fmt.Sprintf("my name is %v", a.name)
}

func (a *Animal) Name() string {
	return a.name
}

// Dog 子类狗
type Dog struct {
	Animal
	Gender string
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and my gender is %v", d.Animal.Speak(), d.Gender)
}

func main() {
	d := Dog{
		Animal: Animal{name: "Hachiko"},
		Gender:  "male",
	}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.Play() // Play() 中调用的是基类 Animal.Speak() 方法，而不是 Dog.Speak()

	a:="test"
	b:="test"
	fmt.Print(&a==&b)
}
