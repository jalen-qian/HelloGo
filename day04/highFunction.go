package main

import (
	"errors"
	"fmt"
)

//高阶函数

//1.函数作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

//2.函数作为返回值
//写一个函数，传入要进行的运算名称"+"、"-"等，返回真正运算的函数
func do(oper string) (func(int, int) int, error) {
	switch oper {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	case "*":
		return mul, nil
	//case "/":
	//	return div, nil

	//或者返回匿名函数
	case "/":
		return func(x, y int) int {
			return x / y
		}, nil
	default:
		return nil, errors.New("操作符" + oper + "不合法")

	}
}

func main() {
	res := calc(10, 20, add)
	fmt.Println(res) // 30

	//参数返回函数
	op, err := do("ss")
	if err == nil {
		fmt.Printf("操作10+15=%d\n", op(10, 15))
	} else {
		fmt.Println(err)
	}

}
