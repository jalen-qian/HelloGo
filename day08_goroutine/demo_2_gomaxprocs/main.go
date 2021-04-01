package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

/**
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数
*/

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A: %d\n", i)
		time.Sleep(time.Millisecond * 100) // 100毫秒
	}
	wg.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B: %d\n", i)
		time.Sleep(time.Millisecond * 100) // 100毫秒
	}
	wg.Done()
}

/**
demo1演示当设置单核运行，但是开启多个goroutine时的运行效果

结论：当单核时，程序会先跑完一个任务，再跑另一个
当多核时，程序会交替跑
B: 0
A: 0
B: 1
A: 1
B: 2
A: 2
A: 3
B: 3
B: 4
A: 4
A: 5
B: 5
B: 6
A: 6
A: 7
B: 7
B: 8
A: 8
A: 9
B: 9
主goroutine运行结束

*/
func demo1() {
	// 设置核心数， 用runtime.GOMAXPROCS()，这里设置单个核心
	runtime.GOMAXPROCS(6)
	wg.Add(2)
	go a()
	go b()

	// 等待 a b都运行完再结束主goroutine
	wg.Wait()
}

func main() {
	demo1()
	fmt.Println("主goroutine运行结束")
}
