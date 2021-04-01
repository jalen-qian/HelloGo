package main

import "fmt"

/**
select 多路复用

有时候我们需要从多个通道中取值，或者发送值给多个通道，但是难保哪个会成功，这时候需要多次判断，类似
a,ok := <- ch1
if ok {...}

b,ok := <- ch2
if ok {...}
...

这样写效率不高，而且代码也很乱，为了解决这个问题，Go语言内置了select关键字，
可以类似于switch一样用多个case来尝试从通道中发送或者接收数据，select会一直等待直到某一个case生效，则执行case下面的语句
1.当只有一个case满足时，执行这个case下的语句
2.当有多个case同时满足时，随机执行一个case下的语句
3.当没有case语句满足时，又没有default，程序会等待，所以一个空的select{}可以用来阻塞main函数
*/
func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case a := <-ch:
			fmt.Println(a)
		case ch <- i:
		default:
		}
	}
}
