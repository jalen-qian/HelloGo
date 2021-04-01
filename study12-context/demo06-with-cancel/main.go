package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	// 没有缓冲区的channel
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case dst <- n:
				time.Sleep(time.Second)
				n++
			case <-ctx.Done():
				return
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
