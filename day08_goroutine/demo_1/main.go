package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/**
goroutine 是Go语言中的一种实现并发的方式，类似于轻量级的线程，
goroutine 是通过Go语言的运行时进行调度和管理的，而线程是操作系统调度和管理的，所以相比于
线程，goroutine消耗的资源更少，更轻量级，并且便于管理。
一个Go语言程序可以同时开启几万甚至十几万个goroutine

goroutine对应一个函数，开启的方式非常简单，在调用函数的时候，前面加上“go”关键字即可
*/

func hello() {
	fmt.Println("hello 我是子goroutine")
	//当子goroutine执行完，waitGroup的计数器减一
	wg.Done()
}

func hello1(i int) {
	fmt.Println("hello 我是子goroutine", i)
	//当子goroutine执行完，waitGroup的计数器减一
	wg.Done()
}

func demo1() {
	/*
		这里发现只打印了hello 我是main，原因是子goroutine还没开始执行，主goroutine就执行完了，在
		Go语言中，如果主goroutine结束，所有的子goroutine都会立刻销毁。

		我们有没有办法让main等待子goroutine呢？
	*/
	//go hello()
	//fmt.Println("hello 我是main")

	//我们可以调用 time.Sleep()，但是这种方式太生硬了，不建议使用
	//正确的做法是，使用 sync.WaitGroup
	wg.Add(1)
	/*
		hello 我是main
		hello 我是子goroutine
	*/
	go hello()
	wg.Wait()
}

/**
  展示通过for循环创造成千上万个goroutine，并且观察执行效果
  结果：打印了十万行，每一行的数字都不一样，且没有按照顺序，杂乱无章，如下：

     hello 我是子goroutine 7199
     hello 我是子goroutine 6950
     hello 我是子goroutine 66269
     hello 我是子goroutine 7073
     hello 我是子goroutine 6703
     ......

  这是因为goroutine是并发执行的，没有先后顺序
*/
func demo2() {
	//直接增加100000计数器
	wg.Add(100000)
	//开启100000个goroutine
	for i := 0; i < 100000; i++ {
		go hello1(i)
	}
	wg.Wait()
}

/**
这个例子中，将hello函数改为匿名函数，看看执行效果

总结：如果这个匿名函数是一个闭包的话，开启goroutine可能会出现问题，因为外部变量的引用是公共的，所以
会出现数据冲突或者重复等等
*/
func demo3() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		//注意，这是一个闭包（匿名函数调用了外部变量）
		//发现问题：打印了很多重复的数字，hello 我是子goroutine 100000这个打印的最多，还有一些其他的
		//原因：因为i是闭包调用了外部的i变量，在子goroutine调用时，i已经增加到了100000，所以打印重复
		//go func() {
		//	fmt.Println("hello 我是子goroutine", i)
		//	wg.Done()
		//}()

		//正确写法，将i作为参数传入匿名函数中，此时打印正确
		go func(i int) {
			fmt.Println("hello 我是子goroutine", i)
			wg.Done()
		}(i)
	}
}

func main() {
	//demo1()

	//demo2()

	demo3()
	fmt.Println("hello 我是main")
}
