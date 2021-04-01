package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	// go语言的运算符
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	// 位运算符 & | ^ << >>
	// 按位与运算&
	var d, e uint = 60, 13
	var f uint
	f = d & e
	fmt.Printf("d & e = %d\n", f) // d & e = 12
	f = d | e
	fmt.Printf("d | e = %d\n", f) // d | e = 61
	f = d ^ e
	fmt.Printf("d ^ e = %d\n", f) // d ^ e = 49

	f = d << 2                     // 左移两位，相当于乘以4
	fmt.Printf("d << 2 = %d\n", f) // d << 2 = 240
	f = d >> 2                     // 左移两位，相当于除以4
	fmt.Printf("d >> 2 = %d\n", f) // d >> 2 = 15
	f = d >> 3                     // 右移三位，相当于除以8
	fmt.Printf("d >> 3 = %d\n", f) // d >> 2 = 7
}
