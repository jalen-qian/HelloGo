package main

import "fmt"

// 结构体的匿名字段
type Person struct {
	string
	int
}

func main() {
	p1 := Person{"张三", 14}
	// p1= main.Person{string:"张三", int:14}
	fmt.Printf("p1= %#v\n", p1)
}
