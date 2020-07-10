package main

import (
	"fmt"
)

var a int = 0 //全局变量
func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//局部变量
	var a, b int
	fmt.Println("请输入任意2个整数：")
	fmt.Scanf("%d%d", &a, &b)
	var max = getMax(a, b)
	fmt.Printf("最大的数是:%d哈哈\n", max)
}
