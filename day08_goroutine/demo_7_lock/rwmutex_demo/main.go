package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写互斥锁：
对于公共资源的竞态问题，分为两种，一种是读资源，一种是写资源，有时候在对公共资源访问时，
读资源的次数比写多得多，这时候如果还是用互斥锁，性能就会相对比较低了，这种场景下，我们
一般用读写互斥锁。

下面的例子里，我们构造读和写两种场景，并假设读资源需要1毫秒，写资源需要10毫秒（一般现实中，读比写
耗时要久），然后开启1000个goroutine去读，开启10个goroutine去写。

我们比较两种情况：
1.全部用互斥锁
2.用读写互斥锁
我们比较这两种写法，最后分别耗时多久
*/

var (
	x      int            //需要操作的全局变量
	wg     sync.WaitGroup //waitGroup
	lock   sync.Mutex     //互斥锁
	rwLock sync.RWMutex   //读写互斥锁
)

/**
读，模拟耗时1毫秒
*/
func read() {
	defer wg.Done()
	//先添加互斥锁
	//lock.Lock()

	//使用读写互斥锁
	rwLock.RLock()
	time.Sleep(time.Millisecond) //模拟消耗1毫秒
	//lock.Unlock()
	rwLock.RUnlock()
}

/**
写，模拟耗时10毫秒
*/
func write() {
	defer wg.Done()
	//使用互斥锁
	//lock.Lock()

	//使用读写互斥锁
	rwLock.Lock()
	x++
	time.Sleep(time.Millisecond * 10) //模拟消耗10毫秒
	//lock.Unlock()
	rwLock.Unlock()
}

func main() {

	//记录程序开始时间
	start := time.Now()
	wg.Add(1010) //计时器增加1010个
	//开启1000个goroutine执行读任务
	for i := 0; i < 1000; i++ {
		go read()
	}

	//开启10个goroutine执行写任务
	for j := 0; j < 10; j++ {
		go write()
	}
	wg.Wait()
	//程序执行完,计算时间间隔
	dur := time.Now().Sub(start)

	/**
	结论：
	1.使用互斥锁：
	执行耗时：1.5960578s

	2.使用读写互斥锁：
	执行耗时：111.7115ms
	可以发现执行效率快了接近10倍
	*/
	fmt.Printf("执行耗时：%v\n", dur)
}
