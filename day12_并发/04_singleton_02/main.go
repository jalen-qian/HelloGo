package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 将03的代码总结一下，就能得到sync.once的实现方式
type once struct {
	mu   sync.Mutex
	done uint32
}

var myOnce = &once{}

func (o *once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	myOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	s := GetInstance()
	fmt.Printf("单例模式的对象：%#v\n", s)
}
