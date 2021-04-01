package main

import "fmt"

/**
类型和接口之间的关系
2.一个接口可以被多个类型实现（多个类型实现同一个接口）
*/

type sayer interface {
	say()
}

type dog struct {
	name string
}

type cat struct {
	name string
}

// 实现同一个接口
func (d dog) say() {
	fmt.Printf("%s在汪汪汪~\n", d.name)
}

func (c cat) say() {
	fmt.Printf("%s在喵喵喵~\n", c.name)
}

/**
实现一个函数，这个函数中传入实现了“叫”接口的变量
这样就可以不管是猫还是狗，只要打，就会“叫”
*/
func hit(s Sayer) {
	s.say()
}

func main() {
	var s sayer
	s = cat{"小花猫"}
	s.say() // 小花猫在喵喵喵~
	s = dog{"小黑狗"}
	s.say() // 小黑狗在汪汪汪~

	hit(cat{"小猫"})
}
