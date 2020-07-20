package main

import (
	"fmt"
	"sync"
)

/**
互斥锁：
有时候多个goroutine同时操作同一资源，会出现竞态问题（数据竞态），这时需要互斥锁来解决此问题

下面的例子，开启两个goroutine同时对全局变量x递增50000次，如果没有竞态问题，最终结果应该是100000
如果有的话，由于互相抢占，会导致最终结果不一定是100000
*/

var (
	x    int
	wg   sync.WaitGroup
	look sync.Mutex //互斥锁
)

/**
这个函数中，不加锁，观察结果
最终递增结果为：73617
C:\Jalen\Programming\GoPath\src\HelloGo\day08_goroutine\demo_7_lock>demo_7_lock.exe
最终递增结果为：79402
C:\Jalen\Programming\GoPath\src\HelloGo\day08_goroutine\demo_7_lock>demo_7_lock.exe
最终递增结果为：91486

结论：每次都不同，且不是预期的结果100000

*/
func add1() {
	for i := 0; i < 50000; i++ {
		x++ //递增
	}
	wg.Done()
}

/**
加一个互斥锁，结果如下：
最终递增结果为：100000

C:\Jalen\Programming\GoPath\src\HelloGo\day08_goroutine\demo_7_lock>demo_7_lock.exe
最终递增结果为：100000

C:\Jalen\Programming\GoPath\src\HelloGo\day08_goroutine\demo_7_lock>demo_7_lock.exe
最终递增结果为：100000

发现不管执行多少次，最终结果都是得到100000
*/
func add2() {
	for i := 0; i < 50000; i++ {
		//操作全局变量之前锁住
		look.Lock()
		x++ //递增
		//操作完解锁
		look.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	//先测试不加锁
	//go add1()
	//go add1()

	//再测试加锁的情况
	go add2()
	go add2()

	wg.Wait()
	//打印最终结果
	fmt.Printf("最终递增结果为：%d\n", x)
}
