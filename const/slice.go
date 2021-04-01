package main

import (
	"fmt"
)

func main() {
	// 切片：slice，可以理解为任意长度的数组，定义时，[]中不用填入长度
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a) // 输出  [1 2 3 4 5]

	// 切片可以通过 “make”函数来创建
	var b []int = make([]int, 3)
	fmt.Println(b) // [0 0 0]

	b[0] = 2
	fmt.Println(b) // [2 0 0]

	// make函数第三个参数
	// 定义切片，长度是5，最大长度是6
	var c []int = make([]int, 5, 6)
	fmt.Println(c) // [0 0 0 0 0]
	// c[5] = 2
	// fmt.Println(c) // 报错，超过引用，后续要学习下如何给切片扩容

	d := [5]int{1, 2, 3, 4, 5} // 定义数组d
	// 定义切片e，是数组d的引用，如果改变了数组d的内容，e也会改变
	var e []int = d[:]

	fmt.Println(e) // [1,2,3,4,5]

	d[4] = 6
	fmt.Println(e) // [1,2,3,4,6]

	/** 从数组的局部创建切片 */
	f := [6]int{1, 2, 3, 4, 5, 6}

	// 从f的第1到第4号位置，创建一个切片，注意，不包括4号位置 1-4，包括1不包括4
	g := f[1:4]

	fmt.Println(g) //[2 3 4]

	// 切片g是数组f局部的一个引用，所以如果改变了f[1] f[2] f[3]的值，切片的值也会跟着变
	f[2] = 100
	fmt.Println(g) //[2 100 4]
	// 同样的，如果要定义从头一直到某个值的切片，则定义方法是 g:= [:4]  g为[1 2 3 4]
	// g:= [2:]  g为[3 4 5]

	h := f[:4]
	fmt.Println(h) //[1 2 100 4]

	/** len()函数和 cap()函数 len()函数获取切片的长度，cap()函数获取切片的最大长度 **/

	i := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(i), cap(i), i) // len=3 cap=5 slice=[0 0 0]

	// 空切片

	var j []int // j是一个空切片

	fmt.Printf("len=%d cap=%d slice=%v\n", len(j), cap(j), j) // len=0 cap=0 slice=[]

	if j == nil { // j是空切片
		fmt.Println("j是空切片")
	} else {
		fmt.Println("j不是空切片")
	}

	/** append 函数和 copy函数 */
	var number []int

	// 1.向
}

// 切片信息输出函数
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
