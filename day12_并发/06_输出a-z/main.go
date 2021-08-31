package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 开启两个协程，依次输出 a - z

func main() {
	c1 := make(chan rune, 1)
	c2 := make(chan rune, 1)
	var once uint32
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if atomic.LoadUint32(&once) == 0 {
			fmt.Printf("协程1输出%c\n", 'a')
			// 告诉协程2，我输出了'a'，该你输出'b'了
			c2 <- 'b'
			atomic.StoreUint32(&once, 1)
		}
	loop:
		for {
			select {
			// 等待协程2通知我输出
			case ch := <-c1:
				if ch <= 'z' {
					fmt.Printf("协程1输出%c\n", ch)
					c2 <- ch + 1
				} else {
					wg.Done()
					break loop
				}
			}
		}
		//
	}()
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			// 等待协程1通知我输出
			case ch := <-c2:
				if ch <= 'z' {
					fmt.Printf("协程2输出%c\n", ch)
					c1 <- ch + 1
				} else {
					wg.Done()
					break loop
				}
			}
		}
	}()
	wg.Wait()
}
