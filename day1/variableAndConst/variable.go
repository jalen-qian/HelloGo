package main

import "fmt"

var name string
var age int
var isOk bool

var (
	a string
	b int
	c float32
)

//d := "hello" // := 只能用于声明局部变量

var m = 100 //声明全局变量m，全局变量声明后可以不使用，局部变量声明后必须使用

func main() {
	n := "abb"
	m := 100 //声明局部变量m，并使用:=

	//100
	//abb
	fmt.Println(m, n)
}
