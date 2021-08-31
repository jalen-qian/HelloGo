package main

import (
	"fmt"
	"sync"
)

var (
	wg    sync.WaitGroup
	total struct {
		sync.Mutex
		value int
	}
)

func worker() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		total.Lock()
		total.value++
		total.Unlock()
	}
}

func main() {
	wg.Add(3)
	go worker()
	go worker()
	go worker()
	wg.Wait()
	fmt.Printf("最终的值是：%v\n", total.value)
}
