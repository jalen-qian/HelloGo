package main

import "fmt"

//定义接口
type Sayer interface {
	say()
}

type Dog struct {}

type Cat struct {}

func (d Dog)say()  {
	fmt.Println("wang~")
}

func (c Cat)say()  {
	fmt.Println("miao~")
}

type Cow struct {

}

//确保Cow一定实现Sayer接口，这里体会“_”的妙用
var _ Sayer = &Cow{} //如果Cow没有实现Sayer的所有方法，那么这一句会编译报错

func (c Cow)say()  {

}

func main() {
	var x Sayer
	d := Dog{}
	c := Cat{}
	d.say()
	c.say()

	x = d
	x.say()
	x = c
	x.say()
}
