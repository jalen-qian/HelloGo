package main

import "fmt"

/**
接口嵌套接口：

接口可以通过嵌套不同的接口，来产生新的接口
*/

//会“叫”的事物，都可以实现sayer接口
type sayer1 interface {
	say()
}

//会“移动”的事物，都可以实现move接口
type mover interface {
	move()
}

//一般动物又会叫，又会动，所以动物接口可以通过嵌套上面两个接口实现
//嵌套的接口使用和普通接口一个样，这里我们创建animal接口
type animal interface {
	sayer1
	mover
}

//创建一个“狗”结构体，来实现动物接口，因为狗是动物
type dog1 struct {
	name string
}

//狗需要实现animal接口
func (d dog1) say() {
	fmt.Printf("%s会叫汪汪汪~\n", d.name)
}

func (d dog1) move() {
	fmt.Printf("%s会跑~\n", d.name)
}

func main() {
	var animal animal = dog1{name: "旺财"}
	animal.move()
	animal.say()
}
