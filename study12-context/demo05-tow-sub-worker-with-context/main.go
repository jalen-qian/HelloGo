package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
方式3：采用context，当子goroutine又启动goroutine时，只需要将context传入即可
*/
var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker1 working...")
		time.Sleep(time.Second)
		select {
		//等待接收上级指示
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func worker2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("worker2 working...")
		time.Sleep(time.Second)
		select {
		//等待接收上级指示
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 10)
	cancel() //调用cancel通知子goroutine任务结束
	wg.Wait()
	fmt.Println("over")
}
