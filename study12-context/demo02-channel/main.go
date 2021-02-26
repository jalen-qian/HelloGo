package main

import (
	"fmt"
	"sync"
	"time"
)

/**
方式1：采用全局变量
缺点：
不易规范和统一，全局变量在跨包调用时无法统一；
如果worker中再启动goroutine，就无法控制了。
*/
var wg sync.WaitGroup
var exit bool

func worker() {
	defer wg.Done()
	for {
		fmt.Println("working...")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func main() {
	wg.Add(1)
	fmt.Println("start working")
	go worker()
	time.Sleep(time.Second * 10)
	exit = true
	wg.Wait()
	fmt.Println("over")
}
