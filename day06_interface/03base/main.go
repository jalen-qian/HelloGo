package main

import "fmt"

/**
类型和接口之间的关系
*/

/* 1.一个类型可以实现多个接口 */
//一条狗既可以叫，又会跑，可以实现两个不同的接口
//两个接口之间相互独立
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Dog struct {
	name string
}

// 实现Sayer接口
func (d Dog) say() {
	fmt.Printf("%s在汪汪汪~\n", d.name)
}

func (d Dog) move() {
	fmt.Printf("%s在跑~\n", d.name)
}

func main() {
	var m Mover = Dog{"旺财"}
	m.move() // 旺财在跑~

	var s Sayer = Dog{"旺财"}
	s.say() // 旺财在汪汪汪~
}
