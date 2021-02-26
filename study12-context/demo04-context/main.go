package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
方式3：官方示例，采用context
*/
var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("working...")
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
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 10)
	cancel() //调用cancel通知子goroutine任务结束
	wg.Wait()
	fmt.Println("over")
}
