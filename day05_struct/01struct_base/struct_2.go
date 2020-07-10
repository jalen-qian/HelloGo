package _1struct_base

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	//结构体指针：使用new关键字创建结构体，分配了内存并得到一个结构体指针
	var s1 = new(Student)
	//s1类型： *main.Student
	fmt.Printf("s1类型： %T\n", s1)

	//new初始化了一个结构体，结构体的属性值默认是各个类型的零值，并返回了一个地址
	//s1的值&main.Student{name:"", age:0}
	fmt.Printf("s1的值%#v\n", s1)

	//在Go语言中，支持直接通过指针“.”来访问结构体的成员
	s1.name = "张三"
	s1.age = 18

	//通过对指针取*，来获取结构体实例
	//main.Student{name:"张三", age:18}
	fmt.Printf("%#v\n", *s1)

	//通过取地址符号来初始化结构体，这个相当于进行了一次new操作
	s2 := &Student{}

	//s2的类型和值：&main.Student{name:"", age:0}
	fmt.Printf("s2的类型和值：%#v\n", s2)

}
