package main

import "fmt"

/**
channel是一种特殊的数据类型，用来在多个goroutine之间进行数据传递和通讯
*/

func main() {
	// channel是一种引用类型，和切片Slice以及Map一样需要用make函数初始化才能使用
	var ch1 chan int
	ch1 = make(chan int, 1) // 初始化通道，并指定通道的缓冲区大小是1

	ch2 := make(chan int) // 定义并初始化通道，没有缓冲区

	// 向通道中发送数据
	ch1 <- 10

	// 从通道中接收数据
	a := <-ch1
	// 10
	fmt.Println(a)

	/**
	  如果向没有缓冲区的通道发送数据，其他goroutine又没有从通道取数据的话，会阻塞
	（这里由于没有其他协程取数据，故而死锁了）
	  没有冲区的通道，必须同时有一个goroutine从这个通道中取出数据，才不会报错，否则会死锁
	  fatal error: all goroutines are asleep - deadlock!
	*/
	//ch2 <- 20

	// 获取通道内数据的个数
	l := len(ch1)
	fmt.Printf("len of ch1 : %d\n", l)
	// 获取通道缓冲区大小
	c := cap(ch2)
	fmt.Printf("cap of ch2 : %d\n", c)

	// 使用内置的close()函数关闭通道
	close(ch1)
	close(ch2)

	// 关闭一个已经关闭的通道，会报panic
	// close(ch1)
}
