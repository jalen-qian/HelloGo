package main

import (
	"fmt"
	"math"
	"reflect"
)

/**
 * 函数的定义
 */
func sum(a int, b int) int {
	return a + b
}

/**
* 函数返回多个参数，注意函数的定义方式
* 给定任意的三个数字，函数按照从小到大的顺序排序后返回
 */
func sort(a int, b int, c int) (int, int, int) {
	if a <= b && b <= c {
		return a, b, c
	}
	if a <= c && c <= b {
		return a, c, b
	}
	if b <= a && a <= c {
		return b, a, c
	}
	if b <= c && c <= a {
		return b, c, a
	}
	if c <= a && a <= b {
		return c, a, b
	}
	if c <= b && b <= a {
		return c, b, a
	}
	return a, b, c
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func typeof1(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	//var a,b int
	//fmt.Println("请输入任意2个整数：")
	//fmt.Scanf("%d%d",&a,&b)
	//sum := sum(a,b)
	//fmt.Printf("%d+%d=%d\n",a,b,sum)

	//排序后的结果为
	var d, e, f = sort(2, 3, 1)
	fmt.Println(d, e, f)

	/*
	 * Go函数作为实参
	 */
	var getSqrt = func(a float64) float64 {
		return math.Sqrt(a)
	}
	//这里可以打印getSqrt变量的类型
	fmt.Println("getSqrt的类型是", typeof1(getSqrt)) // getSqrt的类型是 func(float64) float64
	var x = getSqrt(9)
	fmt.Println(x) //输出3

}
