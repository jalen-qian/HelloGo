package main

import "fmt"

/**
面试题，请问下面的代码执行结果是什么
*/

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	//打印结果：
	//小王子 => 大王八
	//娜扎 => 大王八
	//大王八 => 大王八
	//为什么都是“大王八”？
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	/*
		解释：这里和for range循环的机制有关

		for _, stu := range stus {
			m[stu.name] = &stu
		}
		这里循环的过程如下：
		1.创建一个stu变量，这个变量用来接收每次循环获得的值
		2.当进行循环时，将循环到的stus中的一个结构体赋值给stu（因为结构体是值类型，所以这里是值拷贝）
		3.每次存入到map中的地址是同一个，因为存的是stu变量地址，这个地址没有改变过
		4.最后一次循环，将“大王八”这个结构体赋值给了stu变量，所以map中存储的指针都是指向“大王八”
	*/

	//如何改进，得到我们真正想要的效果？
	//我们可以每次循环都新创建一个变量，这个变量接受stu的值，并将地址存储到map
	for _, stu := range stus {
		mStu := stu
		m[stu.name] = &mStu
	}
	/*
		优化后的结果
		大王八 => 大王八
		小王子 => 小王子
		娜扎 => 娜扎
	*/
	fmt.Println("优化后的结果")
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
