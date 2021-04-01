package main

import (
	"fmt"
	"time"
)

/**
GO处理并发是通过轻量级线程，即“协程”
协程的调度是由GoLang运行时进行管理的，具体开启协程方法如下（使用`go`关键字）:
    go 函数名(参数列表)

Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
同一个程序中的所有 goroutine 共享同一个地址空间
*/

/*函数say,每间隔100毫秒打印一次s字符串，如果同时在两个协程中调用say
  并且打印不同的单词，那么这两个单词应该会交替出现（因为是并发执行）
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 开启一个协程，打印单词world
	go say("world")
	// 同时主协程打印单词hello
	say("hello")
	/*
		可以看到输出如下，且每次运行结果都不一样，说明是异步执行的（并发）
		world
		hello
		world
		hello
		world
		hello
		hello
		world
		hello
		world
	*/

	// 如果main函数这样写：
	// go say("hello")
	// go say("world")
	// 这表示主线程只是开启了两个协程，然后就结束了，所以什么都不会打印，直接结束
}
