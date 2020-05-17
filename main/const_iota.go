package main

import (
	"fmt"
)

func main() {
	//iota 特殊常量，初始值为0，可以在编译时由编译器赋值，iota的值会依次递增
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println("a=", a, ",b=", b, ",c=", c) //a= 0 ,b= 1 ,c= 2

	/*const (
		d = iota
		e = 2
		f = iota
		g = iota
	)
	fmt.Println("d=",d,",e=",e,",f=",f,",g=",g)*/ //d= 0 ,e= 2 ,f= 2 ,g= 3

	/**
	 * 总结，itoa的值，编译时是按照itoa关键字出现在const语句块的位置（可以理解为行索引）
	 */
	const (
		d = iota
		e = 9
		f = iota
		g = 8
		h = iota
	)
	fmt.Println("d=", d, ",e=", e, ",f=", f, ",g=", g, "h=", h) //d= 0 ,e= 9 ,f= 2 ,g= 8 h= 4

}
