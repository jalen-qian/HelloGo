package main

import (
	"fmt"
)

//结构体继承
//可以利用结构体嵌套实现类似于Java语言的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type Dog struct {
	Feet int
	*Animal
}

//构造函数
func NewDog(Feet int, animal *Animal) *Dog {
	return &Dog{
		Feet:   Feet,
		Animal: animal,
	}
}

func (d *Dog) wang() {
	fmt.Printf("%s汪汪汪~\n", d.name)
}

func main() {
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{
			"乐乐",
		},
	}
	//调用了Animal的move，类似于继承
	dog.move()
	//调用自己的method
	dog.wang()

	dog1 := NewDog(4, &Animal{"旺财"})
	dog1.move()
	dog1.wang()

}
