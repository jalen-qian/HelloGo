package main

import "fmt"

func main() {
	var a complex64 = 2 + 3i
	fmt.Println(a) //(2+3i)

	var b complex128 = 3 + 4i
	fmt.Println(b) //(3+4i)

	// 默认为complex128
	c := 4 + 5i
	fmt.Printf("c的类型为%T", c) // c的类型为complex128
}
