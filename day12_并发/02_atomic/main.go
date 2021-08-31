package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 第二种方式实现原子性

var value uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddUint64(&value, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go worker(&wg)
	}
	wg.Wait()
	fmt.Printf("使用atomic包中的函数实现原子性，value：%v\n", value)
}
