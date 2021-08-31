package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 基于atomic的原子性，加上互斥锁实现单例模式
type singleton struct{}

var (
	instance        *singleton
	initializeCount uint32
	mu              sync.Mutex
)

func getInstance() *singleton {
	// 已经
	if atomic.LoadUint32(&initializeCount) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initializeCount, 1)
		instance = &singleton{}
	}

	return instance
}

func main() {
	instance := getInstance()
	fmt.Println(instance)
}
