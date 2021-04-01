package main

import "fmt"

func main() {
	a := 3.24322
	fmt.Printf("a=%f\n", a) // a=3.243220
	// 打印a的类型，说明默认是float64
	fmt.Printf("a的类型是%T\n", a) // a的类型是float64

	var b float32 = 3.22
	fmt.Printf("b=%f\n", b) // b=3.220000

	// 将b赋值给a，报错，不能赋值
	// a = b //cannot use b (type float32) as type float64 in assignment

	// 如果要赋值，使用强制类型转换
	a = float64(b)
	fmt.Printf("a=%f,a的类型是%T\n", a, a) // a=3.220000,a的类型是float64
}
