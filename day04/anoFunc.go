package main

import "fmt"

// 匿名函数和闭包
func main() {
	// 匿名函数保存到变量
	add := func(x, y int) int {
		return x + y
	}
	fmt.Printf("10 + 20 = %d\n", add(10, 20)) // 10 + 20 = 30

	// 匿名函数还可以立即调用
	res := func(x, y int) int {
		return x - y
	}(10, 20)
	fmt.Printf("10 - 20 = %d\n", res) // 10 - 20 = -10

	// 闭包
	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体
}
