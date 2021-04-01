package main

import "fmt"

func fun1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func fun2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func fun3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func fun4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(fun1()) // 5
	fmt.Println(fun2()) // 6
	fmt.Println(fun3()) // 5
	fmt.Println(fun4()) // 5
}
