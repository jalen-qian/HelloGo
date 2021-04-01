package main

import (
	"fmt"
	"strings"
)

// 闭包
// 闭包 = 函数+外部变量的引用

// 1.函数返回一个匿名函数
func a() func() {
	return func() {
		fmt.Println("hello GO")
	}
}

// 2.闭包，如果返回的匿名函数引用了此匿名函数作用域外的变量，这就是一个闭包
func adder(num int) func() int {
	return func() int {
		num++
		return num
	}
}

// 3.闭包进阶，利用闭包写一个自定义的添加文件后缀的工具
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

func main() {
	// sayHello变量就是a()函数返回的匿名函数
	sayHello := a()
	sayHello()

	adder1 := adder(10)
	fmt.Println(adder1()) // 11
	fmt.Println(adder1()) // 12

	// 调用闭包，生成一个自动添加".jpg"后缀的函数
	suffixJpg := makeSuffix(".jpg")
	fmt.Println(suffixJpg("aaa"))     // aaa.jpg
	fmt.Println(suffixJpg("bbb"))     // bbb.jpg
	fmt.Println(suffixJpg("ccc.jpg")) // ccc.jpg

	// 调用闭包，生成一个自动添加".txt"后缀的函数
	suffixTxt := makeSuffix(".txt")
	fmt.Println(suffixTxt("aaa")) // aaa.txt
	fmt.Println(suffixTxt("bbb")) // bbb.txt
}
