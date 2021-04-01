package main

import "fmt"

// defer 延迟执行，defer后面的语句将会在函数执行结束之前执行，并且是倒序的

func testDefer() {
	fmt.Println("start")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	fmt.Println("end")
}

func main() {
	// start
	// end
	// defer3
	// defer2
	// defer1
	testDefer()
}
