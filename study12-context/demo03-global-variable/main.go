package main

import (
	"fmt"
	"sync"
	"time"
)

/**
方式2：采用channel
缺点：

*/
var wg sync.WaitGroup

func worker(exitChan chan struct{}) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("working...")
		time.Sleep(time.Second)
		select {
		// 等待接收上级指示
		case <-exitChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	fmt.Println("start working")
	var exitChan chan struct{}
	exitChan = make(chan struct{}, 1)
	go worker(exitChan)
	time.Sleep(time.Second * 10)
	exitChan <- struct{}{} // 发送消息，结束
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}
