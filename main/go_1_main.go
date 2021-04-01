package main

import (
	"fmt"
)

// 这种写法一般用于声明全局变量
var (
	j int
	k string
)

func main() {
	/*
		const name,age = "钱文军", 25
		fmt.Println(name , "今年", age , "岁了")
		const hello,world = "你好", "世界"
		fmt.Println(hello,world)
		fmt.Println("我是菜鸟，今天第一次写Go")
	*/
	var a string = "我是一个字符串变量"
	fmt.Println(a)

	var b, c int = 1, 2

	fmt.Println(b, c)

	// 不声明 type
	d := "sss"

	fmt.Println(d)

	// 声明了不赋值，默认值为 false
	var e bool

	fmt.Println(e)

	// 声明了不赋值，默认值为0
	var f uint8

	fmt.Println(f)

	// 声明指针类型，不赋值，默认为<nil>
	var g *int
	fmt.Println(g)

	// 通过:=声明新变量
	h, i := 2, "heoo"
	fmt.Println(h, i)

	// 如果:=之前的不是新变量，则会报错  no new variables on left side of :=
	// i := 3
	// fmt.Println(i)

	fmt.Println(j, k)

	// 局部变量声明了，必须被使用，否则编译器报错 l declared and not used
	l := "sssssss"
	fmt.Println("2222", l)

	// 指针类型
	m := &l
	fmt.Println(m) // 0xc0000481f0

	// iota 常量 编译时由编译器赋值，n为0，o为3 p为2
	const (
		n = iota
		o = 3
		p = iota
	)
	fmt.Println(n, o, p)
}
