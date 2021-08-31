package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	tickets = make(chan int, 100)
	wg      sync.WaitGroup
)

func saleTickets(i int) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		// 有余票
		if len(tickets) > 0 {
			// 取出一张余票
			n := <-tickets
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("售出第%d张票\n", n)
		} else {
			// 已售罄
			fmt.Printf("窗口%d已售罄\n", i)
			break
		}
	}
}

func main() {
	// 准备100张票
	for i := 0; i < 100; i++ {
		tickets <- i + 1
	}
	wg.Add(4)
	// 开始售票
	go saleTickets(1)
	go saleTickets(2)
	go saleTickets(3)
	go saleTickets(4)
	wg.Wait()
	// 关闭通道
	close(tickets)
}
