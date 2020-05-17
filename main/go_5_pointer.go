package main

import (
	"fmt"
)

/*
	%b-------------二进制
	%c-------------字符型
	%t--------------布尔型
	%s-------------字符串型
	%f-------------浮点型
	%g------------紧凑型的浮点型
	%d-----------数字型
	%p-----变量的内存地址
	%T-----变量的类型
*/

func main() {
	var a int = 20
	var b *int = &a

	fmt.Printf("变量a的地址     %x\n", &a) //变量a的地址     c00000a0a8
	fmt.Printf("指针b指向的地址 %x\n", b)    //指针b指向的地址 c00000a0a8
	fmt.Printf("根据指针b拿到值 %d\n", *b)   //根据指针b拿到值 20

	var c *int

	fmt.Printf("空指针指向的内存值%x\n", c)         //空指针指向的内存值0
	fmt.Printf("该指针是否是空指针：%t\n", c == nil) //该指针是否是空指针：true

	c = &a
	fmt.Printf("c指针指向的内存值%x\n", c)         //输出 c00000a0a8
	fmt.Printf("该指针是否是空指针：%t\n", c == nil) //false
}
