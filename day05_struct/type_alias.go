package main

import "fmt"

// 自定义类型
type MyInt int

// 类型别名
type mInt = int

func main() {
	var a MyInt
	var b mInt
	fmt.Printf("type of a:%T\n", a) // type of a:main.MyInt  新的类型
	fmt.Printf("type of b:%T\n", b) // type of b:int //b还是int类型，只不过mInt是int的别名
}
