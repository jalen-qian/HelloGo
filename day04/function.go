package main

import "fmt"

//函数
//func 函数名(参数)(返回值){
//
//}

//有参数有返回值
func sum(m int, n int) int {
	return m + n
}

//有参数没有返回值
func f1(m int) {
	fmt.Println(m)
}

//没有参数也没有返回值
func f2() {
	fmt.Println("hello")
}

//有时候返回值参数可以有名字，这时候相当于在函数内部声明了一个变量

func sum2(m int, n int) (ret int) {
	ret = m + n
	return //这时候不需要显示返回变量值
}

//多个返回值
func f3() (a int, b int) {
	a = 2
	b = 3
	return
}

//参数类型简写
func f4(x, y int, m, n string, b1, b2 bool) int {
	return x + y
}

//可变参数
//可变参数必须放在所有参数的最后
//go语言函数中没有默认参数的概念，一切都是显示的，明确的传递
func f5(x int, y ...int) {
	//y的类型是一个切片 1 [2 3 4 5 6]
	//y是一个[]int切片
	fmt.Println(x, y)
}

func main() {
	fmt.Println(sum(2, 3))  //5
	fmt.Println(sum2(2, 3)) //5
	f5(1, 2, 3, 4, 5, 6)
}
