package main

import (
	"fmt"
	"sync"
	"time"
)

/**
问题：主goroutine启动了一个子goroutine worker，如何通知这个子goroutine结束呢？
*/
var wg sync.WaitGroup

func worker() {
	defer wg.Done()
	for {
		fmt.Println("working...")
		time.Sleep(time.Second)
	}
}

func main() {
	wg.Add(1)
	fmt.Println("start working")
	go worker()
	wg.Wait()
	fmt.Println("over")
}
