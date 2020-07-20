package main

import "fmt"

/**
通道和goroutine结合实用案例：
需求分析：
1.创建两个通道，开启两个goroutine
2.一个goroutine负责创建1-100的整数，并放入到通道1中
3.另一个goroutine负责从通道1中取出数，并计算这个数的平方，放到通道2中
4.主程序负责从通道2中取出这些计算好平方的数，并打印
*/

/**
创建1-100的整数，并放入到通道1中
这里使用单向通道，保证f1只能存数据到通道中
*/
func f1(ch1 chan<- int) {
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	//数字都放入通道1后就关闭通道
	close(ch1)
}

/**
从通道1中取出数字，并计算平方后放入通道2中
使用单向通道，保证只能从ch1取数据，只能存数据到ch2
*/
func f2(ch1 <-chan int, ch2 chan<- string) {
	//遍历通道的方法1，通过for无限循环，然后如果取不到值了，就说明通道关闭了，再跳出循环
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		a := fmt.Sprintf("%d^2 = %d\n", i, i*i)
		ch2 <- a
	}
	close(ch2)
}
func main() {
	//1.创建两个通道
	ch1 := make(chan int, 100)
	ch2 := make(chan string, 100)
	//2.开启协程，分别工作
	go f2(ch1, ch2)
	go f1(ch1)
	//3.从通道2中获取数据并打印
	//方法2，使用for range循环，底层会自动判断通道是否关闭
	//如果通道没有关闭，for range循环会无限取值，取完通道中的值后会取出对应类型的零值
	//如果通道是关闭的，for range循环取完值就会退出循环
	for res := range ch2 {
		fmt.Println(res)
	}
}
