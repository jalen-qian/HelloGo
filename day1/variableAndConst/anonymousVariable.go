package main

import "fmt"

/**
 * 返回两个变量，姓名和年龄
 */
func studentMsg(name string, age int) (string, int) {
	return name, age
}

func main() {
	// 这里如果只需要姓名，年龄忽略，可以用匿名变量`_`来接收
	name, _ := studentMsg("张三", 25)
	_, age := studentMsg("李四", 28)
	fmt.Println("姓名是", name)
	fmt.Println("年龄是", age)

	var a int32 = '红'
	fmt.Println(a)
}
