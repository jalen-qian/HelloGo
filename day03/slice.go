package main

import "fmt"

//切片

func main() {
	//切片的声明
	var s1 []int
	fmt.Println(s1) //[]
	//声明了但还没初始化，长度和容量都是0
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1)) //len(s1):0, cap(s1):0
	//判断切片是否初始化
	fmt.Println(s1 == nil) //true
	//切片初始化
	s1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1)) //len(s1):5, cap(s1):5

	fmt.Println(s1 == nil) //false

	//切片声明的同时初始化
	var a []string = []string{"hello", "world"}
	fmt.Println(a)

	// 切片是对数组的封装，切片是引用类型，所以如果将切片赋值给另一个切片，只是将引用拷贝，
	// 两个切）片指向的是同一个数组 （切片底层是数组指针、切片长度、切片容量这三个变量的封装
	var s2 []int = s1
	s2[0] = 2       //改变切片s2
	fmt.Println(s1) //[2 2 3 4 5]

	//切片表达式
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := s1[1:3]
	//输出：sl1:[1 2 3 4 5],len(sl1):5,cap(sl1):5
	fmt.Printf("sl1:%v,len(sl1):%v,cap(sl1):%v\n", sl1, len(sl1), cap(sl1))
	//输出：sl2:[2 3],len(sl2):2,cap(sl2):4
	fmt.Printf("sl2:%v,len(sl2):%v,cap(sl2):%v\n", sl2, len(sl2), cap(sl2))
	//sl2容量是4，说明切片sl2的容量是这样计算的：切片sl2的容量从切片指针指向的第一个元素开始，到最后一个元素结束
	sl2[0] = 88
	//[1 2 3 4 5] [88 3] 切片sl1没改变，说明sl2构建时，底层数组重新拷贝了另一份 sl2不是和sl1指向同一个内存地址了
	fmt.Println(sl1, sl2)

	//切片表达式的上下限不是原切片的长度，而是原切片的容量，比如sl2长度是2，但容量是4，那么可以
	//写 sl3 := sl2[3:4]
	sl3 := sl2[3:4]
	//输出：sl3:[5],len(sl3):1,cap(sl3):1
	fmt.Printf("sl3:%v,len(sl3):%v,cap(sl3):%v\n", sl3, len(sl3), cap(sl3))

}
