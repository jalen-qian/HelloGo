package main

import "fmt"

func main() {
	var a int = 10
	// 输出二进制
	fmt.Printf("a的二进制是%b\n", a) // a的二进制是1010

	// 八进制整数，以0开头
	var b int = 077
	fmt.Printf("b的十进制是%d\n", b) // b的十进制是63
	fmt.Printf("b的八进制是%o\n", b) // b的八进制是77

	// 十六进制整数
	c := 0xff
	fmt.Println(c)               // 255 默认以十进制整数输出
	fmt.Printf("c的十六进制是%x\n", c) // c的十六进制是ff
}
