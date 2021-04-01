package main

import (
	"fmt"
)

func main() {
	const (
		a = 1 << iota // 1左移0位（iota为0）
		b = 3 << iota // 3 11左移1位 110 结果为6
		c             // 相当于 3<<2 1100 结果为12
		d             // 相当于 3<<3 11000 结果为24
	)

	fmt.Println(a, b, c, d) // 1 6 12 24
}
