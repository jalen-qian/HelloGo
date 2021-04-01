package main

import "fmt"

func main() {
	a := '中'
	b := 'a'
	// rune类型实际上是int32
	fmt.Printf("a的类型是 %T, b的类型是%T\n", a, b) // a的类型是 int32, b的类型是int32

	// 这里uint8范围是0-255，而字符'中'超过了uint8的范围，utf-8编码'中'的编码是20013
	// var c uint8 = '中'
	// fmt.Println(c) //constant 20013 overflows uint8

	var c uint32 = '中'
	fmt.Println(c) // 20013

	var d rune = '人'
	fmt.Println(d)        // 20154
	fmt.Printf("%T\n", d) // int32 所以说int32 == rune

	// int32类型变量能直接赋值给rune类型，也进一步说明了rune与int32是同一种类型
	var f int32 = '民'
	d = f
	fmt.Printf("%c\n", d) // 民

	s := "hello沙河"
	// 输出：
	// 104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	// 104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
	for _, r := range s { // rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
