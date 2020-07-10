package main

import "fmt"

func main() {
	//copy函数拷贝切片
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 []int
	copy(s3, s1)
	//这里发现拷贝没成功，这是因为s3是nil切片，没有分配内存
	fmt.Println(s1, s2, s3) //[1 2 3] [1 2 3] []

	//使用make函数创建一个切片，这时候是分配了空间的
	s4 := make([]int, 3, 3)
	copy(s4, s1)
	fmt.Println(s1, s2, s4) //[1 2 3] [1 2 3] [1 2 3]

	//由于是拷贝的，如果改变了s1底层数组的值，不会影响s4
	//但是由于s2和s1底层数组是同一个，所以s2会改变
	s1[0] = 100
	fmt.Println(s1, s2, s4) //[100 2 3] [100 2 3] [1 2 3]
}
