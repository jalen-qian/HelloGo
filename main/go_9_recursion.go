package main

import "fmt"

/**
Go实现递归
func recursion{
   recursion()//函数递归调用自身
}
*/
func main() {
	//利用递归计算阶乘，例如5的阶乘是 1*2*3*4*5，N的阶乘可以表示为N*（N-1的阶乘
	//注意0的阶乘是1
	result := calcFactorial(0)
	fmt.Println(result)           //1
	fmt.Println(calcFactorial(5)) //120

}

/**
由于阶乘的值一般比较大，我们用无符号64位整数来存储
*/
func calcFactorial(n uint64) uint64 {
	if n > 0 {
		return n * calcFactorial(n-1)
	}
	return 1
}
