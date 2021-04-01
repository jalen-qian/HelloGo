package main

import "fmt"

func A() {
	fmt.Println("func A")
}

func B() {
	// recover必须定义在defer关键字之后
	// defer必须在panic之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in func B")
}

func C() {
	fmt.Println("func C")
}

func main() {
	// func A
	// recover in B
	// func C
	A()
	B()
	C()
}
