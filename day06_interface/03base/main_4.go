package main

import "fmt"

/**
空接口：
空接口是指没有定义任何方法的接口，所以任何类型都实现了空接口
声明一个空接口的变量，这个变量可以接收任何类型的值
*/

/**
一个函数的参数是空接口类型，这个参数可以传任意类型
*/
func show(x interface{}) {
	fmt.Printf("type : %T, value: %v\n", x, x)
}

func main() {
	var x interface{}
	x = "hello"
	//type : string, value: hello
	fmt.Printf("type : %T, value: %v\n", x, x)
	x = 10
	//type : int, value: 10
	fmt.Printf("type : %T, value: %v\n", x, x)

	x = true
	// type : bool, value: true
	fmt.Printf("type : %T, value: %v\n", x, x)

	show("sss")
	show(false)

	//利用空接口构造的切片或者接口，可以实现任意值的字典
	student := make(map[string]interface{})
	student["name"] = "小王子"
	student["age"] = 18
	student["married"] = false
	student["class"] = "火箭班"

	//map[string]interface {}{"age":18, "class":"火箭班", "married":false, "name":"小王子"}
	fmt.Printf("%#v\n", student)
}
