package main

import "fmt"

//给基本类型添加方法
//我们只能给自己包中定义的基本类型添加方法，Go内置基本类型由于不在我们的包，所以添加不了
//如果我们就想给基本数据类型添加方法，比如给int8类型添加方法，怎么做到呢？需要用到自定义类型

//这里根据int8定义了一个自定义类型
type MyInt8 int8

//给MyInt8添加方法
func (m MyInt8) sayHi() {
	fmt.Println("Hello I am MyInt8!")
}

func main() {
	var m MyInt8
	m = 10
	//Hello I am MyInt8!
	m.sayHi()

}
